package middleware

import (
	"events/pkg/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func jwtMiddleware(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			ctx.Abort()
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		if token == header {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid format"})
			ctx.Abort()
			return
		}

		claims, err := util.VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Set("userid", claims.UserId)
		ctx.Next()

	}
}
