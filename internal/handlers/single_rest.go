package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (m *Repository) SingleRest(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	rest, err := m.DB.GetRestByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if rest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "The restaurant does not exist"})
		return
	}
	ctx.JSON(http.StatusOK, rest)
}
