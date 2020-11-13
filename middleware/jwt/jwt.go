package jwt

import (
	"net/http"
	"time"
	"wh110api/pak/e"
	"wh110api/pak/jwt"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := com.StrTo(c.GetHeader("Authorization")).String()
		if token == "" {
			code = e.INVALID_PARAMS
		} else {

			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL

			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT

			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"success": false,
				"msg":     e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
