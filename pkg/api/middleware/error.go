package middleware

import (
	"fmt"
	"net/http"
	"todone/pkg/log"
	"todone/pkg/terrors"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()

		// エラーログ出力
		uid, ok := c.Get(AuthCtxKey)
		if !ok {
			log.GetAppLogger().Error(fmt.Sprintf("<error:[Unknown]\n    %+v>", err.Err))
		} else {
			log.GetAppLogger().Error(fmt.Sprintf("<error:[%s]\n    %+v>", uid, err.Err))
		}

		if err != nil {
			var todoneError *terrors.TodoneError
			if ok := xerrors.As(err.Err, &todoneError); ok {
				c.AbortWithStatusJSON(todoneError.ErrorCode, todoneError)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, terrors.Stack(err))
			return
		}
	}
}
