package todo_handler

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
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

func ToggleTodoHandler(ctx *gin.Context) {
	todoId := ctx.Param("id")
	var dbTodo todo.Todo

	res := models.DB.Model(&todo.Todo{}).Where("id = ?", todoId).First(&dbTodo)
	if res.Error != nil {
		log.Fatalf("ERROR: no todo found with that id: %v", res.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "todo not found with that id",
		})
	}

	dbTodo.IsComplete = !dbTodo.IsComplete
	models.DB.Save(&dbTodo)

	cwd, _ := os.Getwd()
	tmpl := template.Must(template.ParseFiles(filepath.Join(cwd, "./views/todo.tmpl")))
	newTodoHTML := renderTodoHTML(tmpl, "todo_element", dbTodo)
	ctx.String(http.StatusOK, newTodoHTML)
}

func DeleteTodoHandler(ctx *gin.Context) {
	todoId := ctx.Param("id")
	var dbTodo todo.Todo

	res := models.DB.Model(&todo.Todo{}).Where("id = ?", todoId).Delete(&dbTodo)
	if res.Error != nil {
		log.Fatalf("ERROR: no todo found with that id: %v", res.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "todo not found with that id",
		})
	}
}

func renderTodoHTML(tmpl *template.Template, tmplName string, todo todo.Todo) string {
	var tplBuffer bytes.Buffer
	tmpl.ExecuteTemplate(&tplBuffer, tmplName, todo)
	return tplBuffer.String()
}
