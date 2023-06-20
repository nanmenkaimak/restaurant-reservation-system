package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

// @Summary Get Menu of Restaurants
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags menu
// @Description get menu of restaurants
// @ID get-menu-restaurants
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Food
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/{rest_id}/menu [get]
func (m *Repository) ShowMenuOfRest(ctx *gin.Context) {
	restID, _ := strconv.Atoi(ctx.Param("rest_id"))

	menu, err := m.DB.GetFoodsByRestID(restID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

// @Summary Add Food to Menu
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags owner
// @Description Add Food to Menu
// @ID create-food-to-menu
// @Accept  json
// @Produce  json
// @Param input body models.Food true "list info"
// @Success 200 {object} models.Food
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/{rest_id}/menu [post]
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
