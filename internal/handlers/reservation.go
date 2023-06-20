package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

// @Summary Get All Reservation of Restaurants
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags owner
// @Description get all reservation of restaurants
// @ID get-all-reservation-restaurants
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Reservations
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/owner/reservations [get]
func (m *Repository) ShowAllReservationsOfRest(ctx *gin.Context) {
	restIDstr := ctx.Query("rest_id")
	var reservations []models.Reservations
	var err error
	ownerID, err := m.checkOwner(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if restIDstr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("restaurant is not defined")})
		return
	}

	restID, _ := strconv.Atoi(restIDstr)

	rest, err := m.DB.GetRestByID(restID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rest.OwnerID != ownerID {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("it is not your restaurant")})
		return
	}

	seats, err := m.DB.GetAllTableOfRest(restID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(seats); i++ {
		reservation, err := m.DB.GetReservationsByTableID(seats[i].ID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reservations = append(reservations, reservation...)
	}
	ctx.JSON(http.StatusOK, reservations)
}

// @Summary Create Reservation
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags reservation
// @Description create Reservation
// @ID create-reservation
// @Accept  json
// @Produce  json
// @Param input body models.Reservations true "list info"
// @Success 200 {object} models.Reservations
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /restaurant/reservations/{seat_id} [post]
func (m *Repository) ReserveTable(ctx *gin.Context) {
	seatID, _ := strconv.Atoi(ctx.Param("seat_id"))

	var newReservation models.Reservations

	if err := ctx.BindJSON(&newReservation); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, ok := ctx.Get("id")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Cannot get ID"})
		return
	}

	customerID := int(id.(float64))

	newReservation.SeatID = seatID
	newReservation.CostumerID = customerID

	err := m.DB.InsertReservation(newReservation)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newReservation)
}
