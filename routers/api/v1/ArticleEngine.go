package v1

import (
	"fmt"
	"go-blog/e"
	"go-blog/models"
	. "go-blog/routers/api/v1/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetArticles(c *gin.Context) {

}

func AddArticle(c *gin.Context) {
	var request AddArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := InvalidParamsResponse()
		c.JSON(http.StatusOK, response)
		return
	}
	response := new(AddArticleResponse)
	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			response.Error = append(response.Error, err.Error())
		}
		response.Code = e.INVALID_PARAMS
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if models.ExistArticleByID(request.ID) {
		response.Code = e.ERROR_EXIST_TAG
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	fmt.Println(ConvertToMap(request))
	models.AddArticle(ConvertToMap(request))
	c.JSON(http.StatusOK, response)
}

func EditArticle(c *gin.Context) {

}

func DeleteARticle(c *gin.Context) {

}
