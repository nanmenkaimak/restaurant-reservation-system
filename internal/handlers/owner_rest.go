package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
	"time"
)

// @Summary Get All Owner Restaurants
// @Security ApiKeyAuth
// @Tags owner
// @Description get all owner restaurants
// @ID get-all-owner-restaurants
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Restaurants
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/owner/{owner_id} [get]
func (m *Repository) AllOwnerRests(ctx *gin.Context) {
	ownerID, _ := strconv.Atoi(ctx.Param("owner_id"))

	rests, err := m.DB.GetAllRestsByOwner(ownerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, rests)
}

// @Summary Create Restaurant
// @Security ApiKeyAuth
// @Tags owner
// @Description create Restaurant
// @ID create-restaurant
// @Accept  json
// @Produce  json
// @Param input body models.Restaurants true "list info"
// @Success 200 {object} models.Restaurants
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/new [post]
func (m *Repository) AddRest(ctx *gin.Context) {
	var newRest models.Restaurants

	newRest.CreatedAt = time.Now()
	newRest.UpdatedAt = time.Now()

	if err := ctx.BindJSON(&newRest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerID, err := m.checkOwner(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRest.OwnerID = ownerID

	err = m.DB.InsertRestaurant(newRest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newRest)
}
