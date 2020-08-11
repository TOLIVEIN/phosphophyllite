package controllers

import (
	"fmt"
	"net/http"
	"phosphophyllite/models"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//AddArticleGet ...
func AddArticleGet(context *gin.Context) {
	isLogin := GetSession(context)
	context.HTML(http.StatusOK, "write_article.html", gin.H{
		"isLogin": isLogin,
	})
}

//AddArticlePost ...
func AddArticlePost(context *gin.Context) {
	session := sessions.Default(context)
	author := session.Get("loginuser").(string)

	title := context.PostForm("title")
	tags := context.PostForm("tags")
	brief := context.PostForm("brief")
	content := context.PostForm("content")
	// author := context.PostForm("author")
	fmt.Printf("title: %s, tags: %s", title, tags)

	article := models.Article{ID: 0, Title: title, Tags: tags, Brief: brief, Content: content, Author: author, Createtime: time.Now().Unix()}
	_, err := models.AddArticle(article)

	response := gin.H{}

	if err == nil {
		response = gin.H{
			"code":    1,
			"message": "ok",
		}
	} else {
		response = gin.H{
			"code":    0,
			"message": "error",
		}
	}

	context.JSON(http.StatusOK, response)
}
