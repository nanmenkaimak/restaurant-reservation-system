package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

func (m *Repository) AllRests(ctx *gin.Context) {
	rest_id := ctx.Query("rest_id")
	var rests []models.Restaurants
	var rest models.Restaurants
	var err error
	if rest_id == "" {
		rests, err = m.DB.GetAllRests()
	} else {
		restID, _ := strconv.Atoi(rest_id)
		rest, err = m.DB.GetRestByID(restID)
		if rest.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "The restaurant does not exist"})
			return
		}
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(rests) == 0 {
		rests = []models.Restaurants{rest}
	}
	ctx.JSON(http.StatusOK, rests)
}
