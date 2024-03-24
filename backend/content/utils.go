package content

import (
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Data: data})
}

func responseFailureWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{Code: http.StatusBadRequest, Data: true, Message: message})
}

func responseSuccessWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Data: true, Message: message})
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

func ToUint64(s string) uint64 {
	number, _ := strconv.ParseUint(s, 10, 64)
	return number
}

func ToBigInt(s string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(s, 10)
	return n
}
