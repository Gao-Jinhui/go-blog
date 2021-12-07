package v1

import (
	"go-blog/e"
	"go-blog/models"
	"go-blog/routers/api/v1/interfaces"
	"go-blog/setting"
	"go-blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	var request interfaces.GetTagsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
	} else {
		response := new(interfaces.GetTagsResponse)
		response.Lists = models.GetTags(util.GetPage(c), setting.PageSize, request)
		response.Total = models.GetTagTotal(request)
		code := e.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": *response,
		})
	}
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
