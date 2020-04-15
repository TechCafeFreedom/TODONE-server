package main

import (
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
		panic(err)
	}

	// DBインスタンスの作成
	dbInstance := mysql.CreateSQLInstance()
	defer dbInstance.Close()

	// トランザクションマネージャーの作成
	masterTxManager := tx.NewDBMasterTxManager(dbInstance)

	// APIインスタンスとmiddlewareインスタンスの作成
	userAPI := InitUserAPI(masterTxManager)
	projectAPI := InitProjectAPI(masterTxManager)
	firebaseClient := middleware.CreateFirebaseInstance(projectAPI)

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

	// userAPI
	firebaseAuth := r.Group("")
	firebaseAuth.Use(firebaseClient.MiddlewareFunc())
	{
		firebaseAuth.POST("/user", userAPI.CreateNewUser)
		firebaseAuth.GET("/user", userAPI.GetUserByPK)
	}
	r.GET("/users", userAPI.GetAllUsers)

	// projectAPI
	firebaseAuth.Use(firebaseClient.MiddlewareFunc())
	{
		firebaseAuth.POST("/project", projectAPI.CreateNewProject)
		firebaseAuth.GET("/user/projects", projectAPI.GetProjectsByUserID)
	}
	r.GET("/project", projectAPI.GetProjectByPK)
	r.GET("/projects", projectAPI.GetAllProjects)

	// port: 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
