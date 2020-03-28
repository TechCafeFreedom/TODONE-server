package middleware

import (
	"os"
	"time"
	"todone/pkg/api/handler"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/model"
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
					"id": v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Id: claims["id"].(string),
			}
		},
		Authenticator: userHandler.Login,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			_, ok := data.(*model.User)
			if ok {
				return true
			}

			return false
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
