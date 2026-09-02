package auth

import (
	"log"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v3"
	"github.com/auth0/go-jwt-middleware/v3/validator"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtValidator *validator.Validator) gin.HandlerFunc {
	// Middleware
	jwt, err := jwtmiddleware.New(
		jwtmiddleware.WithValidator(jwtValidator),
	)
	if err != nil {
		log.Fatal("Error while trying to create jwt validator: ", err)
	}

	return func(ctx *gin.Context) {
		var jwtValid bool
		handler := jwt.CheckJWT(
			// Handler a ejecutar si el JWT es válido
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				jwtValid = true
				ctx.Request = r
			}),
		)

		handler.ServeHTTP(ctx.Writer, ctx.Request)

		if !jwtValid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unathorized",
			})
			return
		}

		ctx.Next()
	}
}

func RequireScope(requiredScope string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwtmiddleware.GetClaims[*validator.ValidatedClaims](
			ctx.Request.Context(),
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})

			return
		}

		customClaims, ok := claims.CustomClaims.(*CustomClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})

			return
		}

		scopes := strings.Fields(customClaims.Scope)

		for _, scope := range scopes {
			if scope == requiredScope {
				ctx.Next()

				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "forbidden",
		})
	}
}
