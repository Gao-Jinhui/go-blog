package jwt

import (
	"go-blog/e"
	. "go-blog/routers/api/v1/interfaces"
	"go-blog/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("admin") != "" {
			c.Next()
		} else {
			token, err := c.Cookie("token")
			response := new(BaseResponse)
			if token == "" || err != nil {
				response.Code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				response.Msg = e.GetMsg(response.Code)
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}
			claims, err := util.ParseToken(token)
			if err != nil {
				response.Code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				response.Msg = e.GetMsg(response.Code)
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}
			if time.Now().Unix() > claims.ExpiresAt {
				response.Code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				response.Msg = e.GetMsg(response.Code)
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}
			c.Next()
		}
	}
}
