package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterGet ...
func RegisterGet(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{
		"title": "注册页",
	})
}
