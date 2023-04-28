package middlewares

import (
	"fmt"
	"go-block-api/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("YO MIDDLE")
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "jwt missing"})
			return
		}
		address, err := auth.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("address", address)
		ctx.Next()
	}
}
