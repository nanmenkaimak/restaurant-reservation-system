package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

// @Summary Get All Restaurants
// @Security ApiKeyAuth
// @Tags restaurants
// @Description get all restaurants
// @ID get-all-restaurants
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Restaurants
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant [get]
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

// @Summary Get Restaurant By Id
// @Security ApiKeyAuth
// @Tags restaurants
// @Description get restaurant by id
// @ID get-restaurant-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "list info"
// @Success 200 {object} models.Restaurants
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/{rest_id} [get]
func (m *Repository) SingleRest(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("rest_id"))

	table, err := m.DB.GetRestByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, table)
}
