package v1

import (
	"go-blog/e"
	"go-blog/models"
	. "go-blog/routers/api/v1/interfaces"
	"go-blog/setting"
	"go-blog/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 获取多个文章标签
func GetTags(c *gin.Context) {
	request := new(GetTagsRequest)
	response := new(GetTagsResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Msg = e.GetMsg(response.Code)
	log.Println(request)
	response.Lists = models.GetTags(util.GetPage(c), setting.PageSize, util.ConvertToMap(*request))
	response.Total = models.GetTagTotal(util.ConvertToMap(*request))
	c.JSON(http.StatusOK, response)
}

// @Summary 新增文章标签
func AddTag(c *gin.Context) {
	request := new(AddTagRequest)
	response := new(AddTagResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if models.ExistTagByName(request.Name) {
		response.Code = e.ERROR_EXIST_TAG
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Msg = e.GetMsg(response.Code)
	models.AddTag(request.Name, request.State, request.CreatedBy)
	c.JSON(http.StatusOK, response)
}

// @Summary 修改文章标签
func UpdateTag(c *gin.Context) {
	request := new(UpdateTagRequest)
	response := new(UpdateTagResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if !models.ExistTagByID(request.ID) {
		response.Code = e.ERROR_NOT_EXIST_TAG
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Msg = e.GetMsg(response.Code)
	models.EditTag(request.ID, util.ConvertToMap(*request))
	c.JSON(http.StatusOK, response)
}

// @Summary 删除文章标签
func DeleteTag(c *gin.Context) {
	request := new(DeleteTagRequest)
	response := new(DeleteTagResponse)
	if response.Code = util.BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if !models.ExistTagByID(request.ID) {
		response.Code = e.ERROR_NOT_EXIST_TAG
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Msg = e.GetMsg(response.Code)
	models.DeleteTag(request.ID)
	c.JSON(http.StatusOK, response)
}
