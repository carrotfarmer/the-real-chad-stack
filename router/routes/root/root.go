package root

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"carrotfarmer/chad-stack/models"
	"carrotfarmer/chad-stack/models/user"
)

// this is the root handler
func Handler(ctx *gin.Context) {
	profile := sessions.Default(ctx).Get("profile")
	mProfile := profile.(map[string]interface{})
	var dbUser []user.User

	res := models.DB.Where(&user.User{
		Email: mProfile["email"].(string),
	}).Find(&dbUser)
	if res.Error != nil {
		log.Printf("ERROR: could not find user for todos: %v", res.Error)
	}

	ctx.HTML(http.StatusOK, "index.tmpl", dbUser[0])
}
