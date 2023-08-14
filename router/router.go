package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"carrotfarmer/chad-stack/auth"
	"carrotfarmer/chad-stack/middleware"
	"carrotfarmer/chad-stack/router/routes/callback"
	"carrotfarmer/chad-stack/router/routes/login"
	"carrotfarmer/chad-stack/router/routes/logout"
	"carrotfarmer/chad-stack/router/routes/profile"
)

func New(auth *auth.Authenticator) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("deez_nuts"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("./views/*.tmpl")

	router.GET("/", func(ctx *gin.Context) {
		profile := sessions.Default(ctx).Get("profile")

		ctx.HTML(http.StatusOK, "index.tmpl", profile)
	})

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/profile", middleware.IsAuthenticated, profile.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
