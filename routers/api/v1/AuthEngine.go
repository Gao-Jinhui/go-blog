package v1

import (
	"go-blog/e"
	"go-blog/models"
	. "go-blog/routers/api/v1/interfaces"
	"go-blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	request := new(GetAuthRequest)
	response := new(GetAuthResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if !models.CheckAuth(request.Username, request.Password) {
		response.Code = e.ERROR_AUTH
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	token, err := util.GenerateToken(request.Username)
	if err != nil {
		response.Code = e.ERROR_TOKEN_GENERATE
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, response)
}
