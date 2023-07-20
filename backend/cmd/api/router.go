package api

import (
	"github.com/gin-gonic/gin"
)

type V1Router struct{}

func v1Router(parent *gin.RouterGroup) {
	router := parent.Group("")
	r := V1Router{}
	router.GET("/", r.helloWorld)
	//router.POST("/deposit", helloWorld)
	//router.GET("/status", helloWorld)
	router.POST("/auth", r.auth)
}

func (v *V1Router) helloWorld(c *gin.Context) {
	responseSuccessWithMessage(c, "Hello")
}
