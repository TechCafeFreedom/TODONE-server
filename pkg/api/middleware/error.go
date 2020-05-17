package middleware

import (
	"fmt"
	"net/http"
	"todone/pkg/terrors"
	"todone/pkg/tlog"

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
			tlog.GetAppLogger().Error(fmt.Sprintf("<error:[Unknown]\n    %+v>", err.Err))
		} else {
			tlog.GetAppLogger().Error(fmt.Sprintf("<error:[%s]\n    %+v>", uid, err.Err))
		}

		// エラーレスポンスの送信
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
