package main

import (
	"log"
	"todone/pkg/infrastructure/gcp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	// DBインスタンスとfireauthのクライアントを取得
	dbInstance, firebaseClient := gcp.CreateInstances()
	defer dbInstance.Close()

	// DBインスタンスの作成
	// dbInstance := mysql.CreateSQLInstance()
	// defer dbInstance.Close()

	// APIインスタンスとmiddlewareインスタンスの作成
	projectAPI := InitProjectAPI(dbInstance)
	// firebaseClient := middleware.CreateFirebaseInstance(projectAPI)

	// CORS対応
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	// connection testAPI
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	// projectAPI
	firebaseAuth := r.Group("")
	firebaseAuth.Use(firebaseClient.MiddlewareFunc())
	{
		firebaseAuth.POST("/project", projectAPI.CreateNewProject)
		firebaseAuth.GET("/user/projects", projectAPI.GetProjectsByUserID)
	}
	r.GET("/project", projectAPI.GetProjectByPK)
	r.GET("/projects", projectAPI.GetAllProjects)

	// port: 8080
	r.Run(":8080")
}
