package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Repository) AllRests(ctx *gin.Context) {
	allRests, err := m.DB.GetAllRests()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allRests)
}
