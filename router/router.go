package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"carrotfarmer/chad-stack/auth"
	"carrotfarmer/chad-stack/middleware"
	"carrotfarmer/chad-stack/router/routes/callback"
	"carrotfarmer/chad-stack/router/routes/login"
	"carrotfarmer/chad-stack/router/routes/logout"
	"carrotfarmer/chad-stack/router/routes/profile"
	"carrotfarmer/chad-stack/router/routes/root"
	todo_handler "carrotfarmer/chad-stack/router/routes/todo"
)

func New(auth *auth.Authenticator) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("deez_nuts"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("./views/*.tmpl")

	router.GET("/", root.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/profile", middleware.IsAuthenticated, profile.Handler)
	router.GET("/logout", logout.Handler)

	router.Group("/todo", middleware.IsAuthenticated)
	{
		router.GET("/")
		router.GET("/:id")
		router.POST("/", todo_handler.CreateTodoHandler)
		router.PATCH("/")
	}

	return router
}
