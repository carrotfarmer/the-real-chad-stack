package callback

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"carrotfarmer/chad-stack/auth"
	"carrotfarmer/chad-stack/models"
	"carrotfarmer/chad-stack/models/todo"
	"carrotfarmer/chad-stack/models/user"
)

// callback handler
func Handler(auth *auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.String(http.StatusBadRequest, "invalid state parameter")
			return
		}

		// exchange authorization code for a token
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusUnauthorized, "failed to exchange authorization code for a token")
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// extract email from profile variable
		email := profile["email"].(string)

		// check if user exists in db, else save
		var dbUser user.User

		res := models.DB.Where(&user.User{
			Email: email,
		}).Find(&dbUser)
		if res.Error != nil {
			log.Fatalf("error while searching for user in DB: %v", res.Error)
		}

		if dbUser.Name == "" {
			models.DB.Create(&user.User{
				Name:    profile["name"].(string),
				Picture: profile["picture"].(string),
				Email:   email,
				Todos:   []todo.Todo{},
			})

			log.Print("successfully added user to database")
		}

		// redirect to logged in page
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
