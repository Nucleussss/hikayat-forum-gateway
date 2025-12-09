package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/Nucleussss/hikayat-forum-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	op := "middleware.AuthMiddleware"
	return gin.HandlerFunc(func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			ctx.Abort() // Stop the request
			return
		}

		tokenString := authHeader[7:] // Extract the token part

		mapClaims, err := utils.ValidateJWTToken(tokenString, os.Getenv("JWT_SECRET"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Error validating JWT token"})
			ctx.Abort() // Stop the request
			return
		}

		userID, ok := (*mapClaims)["user_id"].(string)
		// Check if the user_id is present and of the correct type.
		if !ok {
			log.Printf("%s: %v", op, err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Error get userID from MapClaims"})
			ctx.Abort() // Stop the request
			return
		}

		ctx.Set("user_id", userID)
		ctx.Next() // Continue to the next handler
	})
}
