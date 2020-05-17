package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"todone/pkg/api/request/reqheader"
	"todone/pkg/terrors"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type FirebaseAuth interface {
	MiddlewareFunc() gin.HandlerFunc
}

type firebaseAuth struct {
	client *auth.Client
}

const (
	AuthCtxKey = "AUTHED_UID"
)

func CreateFirebaseInstance() FirebaseAuth {
	ctx := context.Background()

	// get credential of firebase
	var opt option.ClientOption
	gcpClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		// local
		opt = option.WithCredentialsFile("todone-29688-firebase-adminsdk-5y0zs-456e1cfaf4.json")
	} else {
		// gcp
		opt = option.WithCredentialsJSON(getFirebaseCredentialJSON(ctx, gcpClient))
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
	var reqHeader reqheader.Auth
	if err := c.BindHeader(&reqHeader); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	jwtToken := strings.Replace(reqHeader.Authorization, "Bearer ", "", 1)

	// JWT の検証
	authedUserToken, err := fa.client.VerifyIDToken(c, jwtToken)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, terrors.Wrapf(err, http.StatusUnauthorized, "トークンの有効期限が過ぎています。", "token has expired."))
		return
	}
	// contextにuidを格納
	c.Set(AuthCtxKey, authedUserToken.UID)
}

// getFirebaseCredentialJSON firebaseの証明書をjsonで取得
func getFirebaseCredentialJSON(ctx context.Context, client *secretmanager.Client) []byte {
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
