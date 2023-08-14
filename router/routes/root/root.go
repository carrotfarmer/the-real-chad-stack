package root

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"carrotfarmer/chad-stack/models"
	"carrotfarmer/chad-stack/models/user"
)

// this is the root handler
func Handler(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.HTML(http.StatusOK, "index.tmpl", nil)
	}

	profile := sessions.Default(ctx).Get("profile")

	mProfile := profile.(map[string]interface{})
	var dbUser []user.User

	models.DB.Model(&user.User{}).Where("email = ?", mProfile["email"].(string)).Preload("Todos").Find(&dbUser)

	ctx.HTML(http.StatusOK, "index.tmpl", dbUser[0])
}
