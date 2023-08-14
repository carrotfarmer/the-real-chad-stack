package todo_handler

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"carrotfarmer/chad-stack/models"
	"carrotfarmer/chad-stack/models/todo"
	"carrotfarmer/chad-stack/models/user"
)

func CreateTodoHandler(ctx *gin.Context) {
	var dbUser user.User

	// haha cry about it
	email := sessions.Default(ctx).Get("profile").(map[string]interface{})["email"].(string)

	models.DB.Model(&user.User{}).Where("email = ?", email).First(&dbUser)

	newTodo := todo.Todo{
		Text:       ctx.Request.FormValue("todo_text"),
		IsComplete: false,
		UserID:     dbUser.ID,
	}

	if err := models.DB.Create(&newTodo).Error; err != nil {
		log.Fatalf("ERROR: could not create todo: %v", err)
	}

	cwd, _ := os.Getwd()
	tmpl := template.Must(template.ParseFiles(filepath.Join(cwd, "./views/todo.tmpl")))
	tmpl.ExecuteTemplate(ctx.Writer, "todo_element", newTodo)
}
