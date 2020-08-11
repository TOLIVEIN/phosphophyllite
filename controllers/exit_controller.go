package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//ExitGet ...
func ExitGet(context *gin.Context) {
	session := sessions.Default(context)
	session.Delete("loginuser")
	// session.Clear()
	session.Save()

	fmt.Println("delete session...", session.Get("loginuser"))
	context.Redirect(http.StatusMovedPermanently, "/")
}
