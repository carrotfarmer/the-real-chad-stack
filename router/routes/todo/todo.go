package todo_handler

import (
	"carrotfarmer/chad-stack/models"
	"carrotfarmer/chad-stack/models/todo"
	"carrotfarmer/chad-stack/models/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(ctx *gin.Context) {
	var dbUser user.User

	// haha cry about it
	email := sessions.Default(ctx).Get("profile").(map[string]interface{})["email"].(string)

	models.DB.Model(&user.User{}).Where("email = ?", email).Find(&dbUser)

	models.DB.Model(&user.User{}).Where("email = ?", email).UpdateColumn("todos", append(dbUser.Todos, todo.Todo{
		Text:       ctx.Request.FormValue("todo_text"),
		IsComplete: false,
	}))
}
