package gcp

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"todone/pkg/api/middleware"
	"todone/pkg/api/request/reqheader"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// CreateInstances gcpの認証が必要なインスタンスを全て作る
func CreateInstances() (*sql.DB, middleware.FirebaseAuth) {
	var db *sql.DB
	var fireAuth middleware.FirebaseAuth

	opt := option.WithCredentialsFile("todone-server-7767abe40934.json")
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx, opt)
	if err != nil {
		log.Printf("failed to create gcp client: %v", err)
		log.Printf("db connection and fireauth auth local env")
		db = createDBInstanceForLocal()
		fireAuth = createFirebaseInstanceForLocal(ctx, "todone-29688-firebase-adminsdk-5y0zs-43a903b550.json")
		return db, fireAuth
	}
	defer client.Close()

	log.Printf("create instances using gcp client")
	db = createDBInstanceForGCP(ctx, client)
	fireAuth = createFirebaseInstanceForGCP(ctx, client)

	return db, fireAuth
}

// createDBInstanceForLocal 環境変数からMYSQLのconfigを取得
func createDBInstanceForLocal() *sql.DB {
	dbuser := os.Getenv("MYSQL_USER")
	if dbuser == "" {
		dbuser = "root"
	}
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	if dbpassword == "" {
		dbpassword = "password"
	}
	dbhost := os.Getenv("MYSQL_HOST")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbname := "todone"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbname)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// createDBInstanceForGCP GCPからMYSQのconfigを取得
func createDBInstanceForGCP(ctx context.Context, gcpClient *secretmanager.Client) *sql.DB {
	var config mysqlConfig
	err := json.Unmarshal(getSecretData(ctx, gcpClient, "todone-server", "mysql-config"), &config)
	if err != nil {
		log.Panicf("failed to marshal json: %v", err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s", config.User, config.Password, config.Protocol, config.Instance, config.DBName))
	if err != nil {
		log.Panicf("failed to open sql: %v", err)
	}
	return db
}

type mysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db"`
	Instance string `json:"instance"`
}

// createFirebaseinstanceForLocal local環境用のfirebaseインスタンスを作成
func createFirebaseInstanceForLocal(ctx context.Context, keyfile string) middleware.FirebaseAuth {
	opt := option.WithCredentialsFile(keyfile)
	// firebase appの作成
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panic(fmt.Errorf("error initializing firebase app: %v", err))
	}

	// firebase admin clientの作成
	client, err := app.Auth(ctx)
	if err != nil {
		log.Panic(fmt.Errorf("error initialize firebase instance. %v", err))
	}

	return &firebaseAuth{
		client: client,
	}
}

// createFirebaseInstance use gcp service
func createFirebaseInstanceForGCP(ctx context.Context, gcpClient *secretmanager.Client) middleware.FirebaseAuth {

	// secret keyの取得
	opt := option.WithCredentialsJSON(getSecretData(ctx, gcpClient, "todone-server", "fireauth-key"))

	// firebase appの作成
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panic(fmt.Errorf("error initializing firebase app: %v", err))
	}

	// firebase admin clientの作成
	client, err := app.Auth(ctx)
	if err != nil {
		log.Panic(fmt.Errorf("error initialize firebase instance. %v", err))
	}

	return &firebaseAuth{
		client: client,
	}
}

// getSecretData secretmanagerと通信する
func getSecretData(ctx context.Context, client *secretmanager.Client, projectID, secretID string) []byte {
	// requestの作成
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretID),
	}

	// get secret value
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Panicf("failed to access secret version: %v", err)
	}

	return result.Payload.Data
}

func (fa *firebaseAuth) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		fa.middlewareImpl(c)
	}
}

func (fa *firebaseAuth) middlewareImpl(c *gin.Context) {
	// Authorizationヘッダーからjwtトークンを取得
	var reqHeader reqheader.ProjectGet
	if err := c.BindHeader(&reqHeader); err != nil {
		c.Error(err)
	}
	jwtToken := strings.Replace(reqHeader.Authorization, "Bearer ", "", 1)

	// JWT の検証
	authedUserToken, err := fa.client.VerifyIDToken(c, jwtToken)
	if err != nil {
		// JWT が無効なら Handler に進まず別処理
		fmt.Printf("error verifying ID token: %v\n", err)
		c.Error(err)
	}
	// contextにuser_idを格納
	c.Set("AUTHED_USER_ID", authedUserToken.UID)
}

type firebaseAuth struct {
	client *auth.Client
}
