package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Repository) AllRests(ctx *gin.Context) {
	allRests, err := m.DB.GetAllRests()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ctx.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, allRests)
}
