package handlers

import (
	"github.com/gin-gonic/gin"
)

func (m *Repository) Home(ctx *gin.Context) {
	ctx.Writer.Write([]byte("hello"))
}
