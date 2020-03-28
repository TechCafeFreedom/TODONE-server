package middleware

import (
	"os"
	"time"
	model "todone/db/mysql/models"
	"todone/pkg/api/handler"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func CreateJwtInstance(userHandler handler.UserHandler) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "todone zone",
		Key:         []byte(os.Getenv("JWT_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"id": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				ID: claims["id"].(string),
			}
		},
		Authenticator: userHandler.Login,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			_, ok := data.(*model.User)
			return ok
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

}
