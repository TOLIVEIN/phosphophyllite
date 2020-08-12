package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//GetSession ...
func GetSession(context *gin.Context) bool {
	session := sessions.Default(context)
	loginuser := session.Get("loginuser")
	fmt.Println("loginuser: ", loginuser)
	if loginuser != nil {
		return true
	}
	return false
}
