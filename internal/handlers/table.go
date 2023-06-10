package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

func (m *Repository) AddTable(ctx *gin.Context) {
	var newTable models.Seats

	if err := ctx.BindJSON(&newTable); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restID, _ := strconv.Atoi(ctx.Param("rest_id"))

	newTable.RestaurantID = restID

	err := m.DB.InsertTable(newTable)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newTable)
}

func (m *Repository) ShowAllTableOfRest(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("rest_id"))

	tables, err := m.DB.GetAllTableOfRest(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tables)
}
