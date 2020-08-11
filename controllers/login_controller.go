package controllers

import (
	"fmt"
	"net/http"
	"phosphophyllite/models"
	"phosphophyllite/utils"

	"github.com/gin-gonic/gin"
)

//LoginGet ...
func LoginGet(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录页",
	})
}

//LoginPost ...
func LoginPost(context *gin.Context) {
	username, password := context.PostForm("username"), context.PostForm("password")
	fmt.Printf("username: %s, password: %s", username, password)

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id: ", id)
	if id > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "登录成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "登录失败",
		})
	}
}
