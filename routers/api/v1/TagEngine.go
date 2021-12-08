package v1

import (
	"go-blog/e"
	"go-blog/models"
	"go-blog/routers/api/v1/interfaces"
	"go-blog/setting"
	"go-blog/util"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	var request interfaces.GetTagsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := InvalidParamsResponse()
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := new(interfaces.GetTagsResponse)
		filter := ConvertToMap(request)
		response.Lists = models.GetTags(util.GetPage(c), setting.PageSize, filter)
		response.Total = models.GetTagTotal(filter)
		response.Code = e.SUCCESS
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
}

//新增文章标签
func AddTag(c *gin.Context) {
	var request interfaces.AddTagRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := InvalidParamsResponse()
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := new(interfaces.AddTagResponse)
		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				response.Error = append(response.Error, err.Error())
			}
			response.Code = e.INVALID_PARAMS
			response.Msg = e.GetMsg(response.Code)
			c.JSON(http.StatusOK, response)
			return
		} else {
			if models.ExistTagByName(request.Name) {
				response.Code = e.ERROR_EXIST_TAG
			} else {
				response.Code = e.SUCCESS
				models.AddTag(request.Name, request.State, request.CreatedBy)
			}
			response.Msg = e.GetMsg(response.Code)
			c.JSON(http.StatusOK, response)
			return
		}
	}
}

//修改文章标签
func EditTag(c *gin.Context) {
	var request interfaces.EditTagRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := InvalidParamsResponse()
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := new(interfaces.EditTagResponse)
		validate := validator.New()
		if err := validate.Struct(&request); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				response.Error = append(response.Error, err.Error())
			}
			response.Code = e.INVALID_PARAMS
			return
		} else {
			if !models.ExistTagByID(request.ID) {
				response.Code = e.ERROR_NOT_EXIST_TAG
			} else {
				response.Code = e.SUCCESS
				data := ConvertToMap(request)
				models.EditTag(request.ID, data)
			}
		}
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
	}
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	var request interfaces.DeleteTagRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := InvalidParamsResponse()
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := new(interfaces.DeleteTagResponse)
		validate := validator.New()
		if err := validate.Struct(&request); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				response.Error = append(response.Error, err.Error())
			}
			response.Code = e.INVALID_PARAMS
			return
		} else {
			if !models.ExistTagByID(request.ID) {
				response.Code = e.ERROR_NOT_EXIST_TAG
			} else {
				response.Code = e.SUCCESS
				models.DeleteTag(request.ID)
			}
		}
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
	}
}

func InvalidParamsResponse() *interfaces.BaseResponse {
	res := new(interfaces.BaseResponse)
	res.Code = e.INVALID_PARAMS
	res.Msg = e.GetMsg(res.Code)
	return res
}

func ConvertToMap(request interface{}) map[string]interface{} {
	typeOfRequest := reflect.TypeOf(request)
	valueOfRequest := reflect.ValueOf(request)
	filter := make(map[string]interface{})
	for pos := 0; pos < typeOfRequest.NumField(); pos++ {
		filter[typeOfRequest.Field(pos).Name] = valueOfRequest.Field(pos).Interface()
	}
	return filter
}
