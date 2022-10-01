package validation

import (
	"fmt"
	"strings"
	"web-chat/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func jsonKeyBuidler(key string) string {
	if len(key) == 0 {
		return ""
	}
	return strings.ToLower(key[:1]) + key[1:]
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required field", jsonKeyBuidler(fe.Field()))
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", jsonKeyBuidler(fe.Field()), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", jsonKeyBuidler(fe.Field()), fe.Param())
	case "email":
		return "invalid email address"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", jsonKeyBuidler(fe.Field()), fe.Param())
	}
	return fmt.Sprintf("%s is not valid", jsonKeyBuidler(fe.Field()))
}

func Bind(ctx *gin.Context, req interface{}) bool {

	if ctx.ContentType() != "application/json" {

		msg := "request body type is not valid"

		response.BadRequest(ctx, msg)
		return false
	}

	if err := ctx.ShouldBind(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []string
			for _, fe := range errs {
				invalidArgs = append(invalidArgs, getErrorMessage(fe))
			}

			response.BadRequest(ctx, invalidArgs)
			return false

		}
		response.InternalServerError(ctx, err.Error())
		return false

	}
	return true
}

func BindUri(ctx *gin.Context, req interface{}) bool {

	if err := ctx.ShouldBindUri(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []string
			for _, fe := range errs {
				invalidArgs = append(invalidArgs, getErrorMessage(fe))
			}

			response.BadRequest(ctx, invalidArgs)
			return false

		}

		response.InternalServerError(ctx, err.Error())
		return false

	}
	return true
}

func BindWith(ctx *gin.Context, req interface{}) bool {

	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []string
			for _, fe := range errs {
				invalidArgs = append(invalidArgs, getErrorMessage(fe))
			}

			response.BadRequest(ctx, invalidArgs)
			return false

		}

		response.InternalServerError(ctx, err.Error())
		return false

	}
	return true
}
