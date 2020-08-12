package controllers

import (
	"net/http"
	"phosphophyllite/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HomeGet ...
func HomeGet(context *gin.Context) {
	isLogin := GetSession(context)

	page, _ := strconv.Atoi(context.Query("page"))
	if page <= 0 {
		page = 1
	}
	var articleList []models.Article
	articleList, _ = models.FindArticleWithPage(page)
	html := models.MakeHomeBlocks(articleList, isLogin)

	homeFooterPage := models.SetHomeFooterPage(page)
	context.HTML(http.StatusOK, "home.html", gin.H{
		// "title":   "首页",
		"isLogin":   isLogin,
		"content":   html,
		"hasFooter": true,
		"page":      homeFooterPage,
	})
}
