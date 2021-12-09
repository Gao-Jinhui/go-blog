package v1

import (
	"fmt"
	"go-blog/e"
	"go-blog/models"
	. "go-blog/routers/api/v1/interfaces"
	"go-blog/setting"
	"go-blog/util"
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
	response.Data = models.GetArticle(*(request.ID))
	c.JSON(http.StatusOK, response)
}

func GetArticlesByTag(c *gin.Context) {
	var request GetArticlesByTagRequest
	response := new(GetArticlesByTagResponse)
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Code = e.INVALID_PARAMS
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		response.Code = e.ERROR_VALIDATOR
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = e.SUCCESS
	response.Msg = e.GetMsg(response.Code)
	response.Data = models.GetArticles(util.GetPage(c), setting.PageSize, ConvertToMap(request))
	c.JSON(http.StatusOK, response)
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
	if !models.ExistArticleByID(request.Tag_ID) {
		response.Code = e.ERROR_NOT_EXIST_TAG
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

func UpdateArticle(c *gin.Context) {
	request := new(UpdateArticleRequest)
	response := new(UpdateArticleResponse)
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Code = e.INVALID_PARAMS
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		response.Code = e.ERROR_VALIDATOR
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
	if response.Code = BindAndValidateParams(c, request); response.Code != e.SUCCESS {
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Msg = e.GetMsg(response.Code)
	models.DeleteArticle(request.ID)
	c.JSON(http.StatusOK, response)
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

func BindAndValidateParams(c *gin.Context, request interface{}) int {
	if err := c.ShouldBindJSON(request); err != nil {
		return e.INVALID_PARAMS
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return e.ERROR_VALIDATOR
	}
	return e.SUCCESS
}
