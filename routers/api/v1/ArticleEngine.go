package v1

import (
	"go-blog/e"
	"go-blog/models"
	. "go-blog/routers/api/v1/interfaces"
	"go-blog/setting"
	"go-blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticleByID(c *gin.Context) {
	request := new(GetArticleByIDRequest)
	response := new(GetArticleByIDResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	response.Data = models.GetArticle(*(request.ID))
	c.JSON(http.StatusOK, response)
}

func GetArticlesByTag(c *gin.Context) {
	request := new(GetArticlesByTagRequest)
	response := new(GetArticlesByTagResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	response.Data = models.GetArticles(util.GetPage(c), setting.PageSize, util.ConvertToMap(*request))
	c.JSON(http.StatusOK, response)
}

func AddArticle(c *gin.Context) {
	request := new(AddArticleRequest)
	response := new(AddArticleResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if !models.ExistTagByID(request.Tag_ID) {
		response.Code = e.ERROR_NOT_EXIST_TAG
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	models.AddArticle(util.ConvertToMap(*request))
	c.JSON(http.StatusOK, response)
}

func UpdateArticle(c *gin.Context) {
	request := new(UpdateArticleRequest)
	response := new(UpdateArticleResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if !models.ExistArticleByID(request.ID) {
		response.Code = e.ERROR_NOT_EXIST_ARTICLE
		response.Msg = e.GetMsg(response.Code)
		return
	}
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	models.UpdateArticle(request.ID, request)
	c.JSON(http.StatusOK, response)
}

func DeleteArticle(c *gin.Context) {
	request := new(DeleteArticleRequest)
	response := new(DeleteArticleResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Msg = e.GetMsg(response.Code)
	models.DeleteArticle(request.ID)
	c.JSON(http.StatusOK, response)
}
