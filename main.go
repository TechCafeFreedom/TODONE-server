package main

import (
	"todone/db/mysql"
	"todone/pkg/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbInstance := mysql.CreateSQLInstance()
	defer dbInstance.Close()

	userAPI := InitUserAPI(dbInstance)

	jwtAuth, err := middleware.CreateJwtInstance(userAPI)
	if err != nil {
		panic(err)
	}

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
	r.POST("/login", jwtAuth.LoginHandler)
	authByJwt := r.Group("")
	authByJwt.Use(jwtAuth.MiddlewareFunc())
	{
		authByJwt.GET("/user/get", userAPI.GetUser)
	}
	r.POST("/user/create", userAPI.CreateNewUser)
	r.GET("/users", userAPI.GetAllUsers)

	// ポートを設定しています。
	r.Run(":8080")
}
