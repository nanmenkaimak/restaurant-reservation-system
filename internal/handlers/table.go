package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

// @Summary Create Table
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags owner
// @Description create Table
// @ID create-table
// @Accept  json
// @Produce  json
// @Param input body models.Seats true "list info"
// @Success 200 {object} models.Seats
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/{rest_id}/table [post]
func (m *Repository) AddTable(ctx *gin.Context) {
	var newTable models.Seats

	if err := ctx.BindJSON(&newTable); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerID, err := m.checkOwner(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restID, _ := strconv.Atoi(ctx.Param("rest_id"))

	rest, err := m.DB.GetRestByID(restID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("cannot get restaurant from database")})
		return
	}

	if rest.OwnerID != ownerID {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "it is not your restaurant"})
		return
	}

	newTable.RestaurantID = restID

	err = m.DB.InsertTable(newTable)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newTable)
}

// @Summary Get All Table of Restaurants
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags table
// @Description get all table of restaurants
// @ID get-all-table-restaurants
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Seats
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/{rest_id}/table [get]
func (m *Repository) ShowAllTableOfRest(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("rest_id"))

	tables, err := m.DB.GetAllTableOfRest(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tables)
}
