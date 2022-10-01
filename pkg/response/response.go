package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Success struct {
	Status   int         `json:"status"`
	Message  interface{} `json:"message"`
	Response interface{} `json:"response"`
}

type Error struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}

type Health struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Health  interface{} `json:"health"`
}

func OK(ctx *gin.Context, result interface{}) {
	ctx.JSON(http.StatusOK, Success{
		Status:   http.StatusOK,
		Message:  http.StatusText(http.StatusOK),
		Response: result,
	})
}

func Created(ctx *gin.Context, result interface{}) {
	ctx.JSON(http.StatusCreated, Success{
		Status:   http.StatusCreated,
		Message:  http.StatusText(http.StatusCreated),
		Response: result,
	})
}

func BadRequest(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusBadRequest, Error{
		Status:  http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Error:   err,
	})
}

func Unauthorized(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusUnauthorized, Error{
		Status:  http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
		Error:   err,
	})
}

func Forbidden(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusForbidden, Error{
		Status:  http.StatusForbidden,
		Message: http.StatusText(http.StatusForbidden),
		Error:   err,
	})
}

func NotFound(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusNotFound, Error{
		Status:  http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
		Error:   err,
	})
}

func InternalServerError(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusInternalServerError, Error{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Error:   err,
	})
}
