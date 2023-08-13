package user

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// login page handler
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	fmt.Println(profile)
	ctx.HTML(http.StatusOK, "user.html", profile)
}
