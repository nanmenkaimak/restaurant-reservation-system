package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

func (m *Repository) ShowMenuOfRest(ctx *gin.Context) {
	restID, _ := strconv.Atoi(ctx.Param("rest_id"))

	menu, err := m.DB.GetFoodsByRestID(restID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

func (m *Repository) AddFoodToMenu(ctx *gin.Context) {
	restID, _ := strconv.Atoi(ctx.Param("rest_id"))

	var newFood models.Food

	if err := ctx.BindJSON(&newFood); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerID, err := m.checkOwner(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rest, err := m.DB.GetRestByID(restID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("cannot get restaurant from database")})
		return
	}

	if rest.OwnerID != ownerID {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "it is not your restaurant"})
		return
	}

	newFood.RestaurantID = restID

	err = m.DB.InsertFood(newFood)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newFood)
}
