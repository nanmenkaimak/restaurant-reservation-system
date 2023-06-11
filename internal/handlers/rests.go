package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

func (m *Repository) Rests(ctx *gin.Context) {
	city := ctx.Query("city")
	var rests []models.Restaurants
	var err error
	if city == "" {
		rests, err = m.DB.GetAllRests()
	} else {
		rests, err = m.DB.GetRestsByCity(city)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rests)
}

func (m *Repository) SingleRest(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("rest_id"))

	table, err := m.DB.GetRestByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, table)
}
