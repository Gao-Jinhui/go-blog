package util

import (
	"go-blog/e"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
