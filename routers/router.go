package routers

import (
	v1 "go-blog/routers/api/v1"
	"go-blog/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/v1")
	{
		apiv1.POST("/tags/get", v1.GetTags)
		apiv1.POST("/tags/add", v1.AddTag)
		apiv1.POST("/tags/update", v1.UpdateTag)
		apiv1.POST("/tags/delete", v1.DeleteTag)
		apiv1.POST("/article/getbyid", v1.GetArticleByID)
		apiv1.POST("/article/getbytag", v1.GetArticlesByTag)
		apiv1.POST("/article/add", v1.AddArticle)
		apiv1.POST("/article/update", v1.UpdateArticle)
		apiv1.POST("/article/delete", v1.DeleteArticle)
	}

	return r
}
