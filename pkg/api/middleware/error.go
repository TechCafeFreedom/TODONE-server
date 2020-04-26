package middleware

import (
	"errors"
	"net/http"
	"todone/pkg/terrors"

	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			var todoneError *terrors.TodoneError
			if ok := errors.As(err.Err, &todoneError); ok {
				c.AbortWithStatusJSON(todoneError.ErrorCode, todoneError)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}
}
