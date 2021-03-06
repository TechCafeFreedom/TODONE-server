package main

import (
	"log"
	"todone/db/mysql"
	"todone/pkg/api/middleware"
	tx "todone/pkg/infrastructure/mysql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	// DBインスタンスの作成
	dbInstance := mysql.CreateSQLInstance()
	defer dbInstance.Close()

	// トランザクションマネージャーの作成
	masterTxManager := tx.NewDBMasterTxManager(dbInstance)

	// APIインスタンスの作成
	userAPI := InitUserAPI(masterTxManager)
	boardAPI := InitBoardAPI(masterTxManager)

	// firebase middlewareの作成
	firebaseClient := middleware.CreateFirebaseInstance()

	// CORS対応
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	// エラー出力用ミドルウェアの使用
	r.Use(middleware.ErrorHandling())

	// connection testAPI
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	// userAPI
	firebaseAuth := r.Group("")
	firebaseAuth.Use(firebaseClient.MiddlewareFunc())
	{
		firebaseAuth.POST("/users/self", userAPI.CreateNewUser)
		firebaseAuth.GET("/users/self", userAPI.GetUserProfile)
	}

	// boardAPI
	firebaseAuth.Use(firebaseClient.MiddlewareFunc())
	{
		firebaseAuth.POST("/boards", boardAPI.CreateNewBoard)
		firebaseAuth.GET("/boards", boardAPI.GetUserBoards)
		firebaseAuth.GET("/boards/:id", boardAPI.GetBoardDetail)
	}

	// port: 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
