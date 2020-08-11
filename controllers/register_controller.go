package controllers

import (
	"net/http"
	"phosphophyllite/models"
	"phosphophyllite/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterGet ...
func RegisterGet(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{
		"title": "注册页",
	})
}

//RegisterPost ...
func RegisterPost(context *gin.Context) {
	username, password, repassword := context.PostForm("username"), context.PostForm("password"), context.PostForm("repassword")
	println(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	println("id: ", id)
	if id > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "用户名已存在",
		})
		return
	}

	password = utils.MD5(password)
	println("MD5 password: ", password)
	user := models.User{ID: 0, Username: username, Password: password, Status: 0, Createtime: time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "注册失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "注册成功",
		})
	}
}

//LoginGet ...
func LoginGet(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录页",
	})
}
