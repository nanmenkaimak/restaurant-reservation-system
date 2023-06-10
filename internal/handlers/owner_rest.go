package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"time"
)

func (m *Repository) AllOwnerRest(ctx *gin.Context) {

}

func (m *Repository) AddRest(ctx *gin.Context) {
	var newRest models.Restaurants

	newRest.CreatedAt = time.Now()
	newRest.UpdatedAt = time.Now()

	if err := ctx.BindJSON(&newRest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerID, ok := ctx.Get("id")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Cannot get ID"})
		return
	}

	newRest.OwnerID = int(ownerID.(float64))

	owner, err := m.DB.GetUserByID(newRest.OwnerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if owner.RoleID != 2 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Only owner can add restaurants"})
		return
	}

	err = m.DB.InsertRestaurant(newRest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newRest)
}
