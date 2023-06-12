package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
	"time"
)

func (m *Repository) AllOwnerRests(ctx *gin.Context) {
	ownerID, _ := strconv.Atoi(ctx.Param("owner_id"))

	rests, err := m.DB.GetAllRestsByOwner(ownerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, rests)
}

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
