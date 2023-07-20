package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	argsAddr = "addr"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Data: data})
}

func responseSuccessWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{Data: true, Message: message})
}

func responseErrUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{Code: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized)})
}

func responseErrValidation(c *gin.Context, message string) {
	c.JSON(http.StatusUnprocessableEntity, Response{Code: http.StatusUnprocessableEntity, Message: message})
}

func responseErrInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)})
}

func responseErrInternalServerErrorWithDetail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{Code: http.StatusInternalServerError, Message: message})
}

func responseErrNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{Code: http.StatusNotFound, Message: http.StatusText(http.StatusNotFound)})
}