package middleware

import (
	"context"
	"fmt"
	"log"
	"strings"
	"todone/pkg/api/handler/project"
	"todone/pkg/api/request/reqheader"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseAuth interface {
	MiddlewareFunc() gin.HandlerFunc
}

type firebaseAuth struct {
	server project.Server
	client *auth.Client
}

func CreateFirebaseInstance(server project.Server) FirebaseAuth {
	// firebase appの作成
	ctx := context.Background()
	opt := option.WithCredentialsFile("todone-29688-firebase-adminsdk-5y0zs-456e1cfaf4.json")
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
