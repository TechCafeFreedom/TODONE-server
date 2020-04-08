package middleware

import (
	"context"
	"fmt"
	"log"
	"strings"
	"todone/pkg/api/handler/project"
	"todone/pkg/api/request/reqheader"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type FirebaseAuth interface {
	MiddlewareFunc() gin.HandlerFunc
}

type firebaseAuth struct {
	server project.Server
	client *auth.Client
}

func CreateFirebaseInstance(server project.Server) FirebaseAuth {

	ctx := context.Background()

	// get credential of firebase
	var opt option.ClientOption
	gcpClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		// local
		opt = option.WithCredentialsFile("todone-29688-firebase-adminsdk-5y0zs-456e1cfaf4.json")
		// opt = option.WithCredentialsFile("todone-29688-firebase-adminsdk-5y0zs-43a903b550.json")
	} else {
		// gcp
		opt = option.WithCredentialsJSON(getFirebaseCredentialJSON(gcpClient, ctx))
	}

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
		server: server,
		client: client,
	}
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

// getFirebaseCredentialJSON
func getFirebaseCredentialJSON(client *secretmanager.Client, ctx context.Context) []byte {
	projectID := "todone-server"
	secretID := "fireauth-key"
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
