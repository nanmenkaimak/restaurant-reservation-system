package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/JWT"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth header"})
			return
		}
		headerInSlice := strings.Split(header, " ")
		if len(headerInSlice) != 2 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong header"})
			return
		}
		claims, err := JWT.ParseToken(headerInSlice[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("id", claims["id"])
		ctx.Next()
	}
}
