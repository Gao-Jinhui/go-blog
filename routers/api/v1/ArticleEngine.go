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

func GetArticleByID(c *gin.Context) {
	var request GetArticleByIDRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := GetInvalidParamsResponse()
		c.JSON(http.StatusOK, response)
		return
	}
	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		c.JSON(http.StatusOK, GetValidatorErrorResponse(err))
		return
	}
	response := new(GetArticleByIDResponse)
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	response.Data = models.GetArticle(request.ID)
	c.JSON(http.StatusOK, response)
}

func GetArticlesByTag(c *gin.Context) {

}

func AddArticle(c *gin.Context) {
	var request AddArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, GetInvalidParamsResponse())
		return
	}
	response := new(AddArticleResponse)
	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		c.JSON(http.StatusOK, GetValidatorErrorResponse(err))
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

func GetValidatorErrorResponse(err error) *ValidatorErrorResponse {
	var response ValidatorErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		response.Error = append(response.Error, err.Error())
	}
	response.Code = e.ERROR_VALIDATOR
	response.Msg = e.GetMsg(response.Code)
	return &response
}
