package todo_handler

import (
	"carrotfarmer/chad-stack/models"
	"carrotfarmer/chad-stack/models/todo"
	"carrotfarmer/chad-stack/models/user"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(ctx *gin.Context) {
	var dbUser user.User

	// haha cry about it
	email := sessions.Default(ctx).Get("profile").(map[string]interface{})["email"].(string)

	models.DB.Model(&user.User{}).Where("email = ?", email).Find(&dbUser)

	fmt.Println("the value should be below this:")
	fmt.Println(ctx.Request.FormValue("todo_text"))

	// create a new todos array with new todo
	// newTodos := append(dbUser.Todos, todo.Todo{
	// 	Text:       ctx.Request.FormValue("todo_text"),
	// 	IsComplete: false,
	// })

	// update todos array in user
	// res := models.DB.Model(&user.User{}).Where("email = ?", email).UpdateColumn("todos", newTodos)
	// if res.Error != nil {
	// 	log.Fatalf("ERROR: could not create todo: %v", res.Error)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": res.Error,
	// 	})
	// }

	models.DB.Model(&user.User{}).Where("email = ?", email).Association("Todos").Append(&todo.Todo{
		Text:       ctx.Request.FormValue("todo_text"),
		IsComplete: false,
		UserId:     dbUser.ID,
	})
}
