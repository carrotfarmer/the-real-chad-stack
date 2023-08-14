package profile

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	ctx.HTML(http.StatusOK, "profile.tmpl", profile)
}
