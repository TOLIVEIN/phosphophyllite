package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//IndexGet ...
func IndexGet(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "首页",
	})
}
