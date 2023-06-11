package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"strconv"
)

func (m *Repository) ShowAllReservationsOfRest(ctx *gin.Context) {
	restIDstr := ctx.Query("rest_id")
	var reservations []models.Reservations
	var err error
	ownerID, err := m.checkForOwner(ctx)
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "it is not your restaurant"})
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

func (m *Repository) checkForOwner(ctx *gin.Context) (int, error) {
	ownerIDstr, ok := ctx.Get("id")
	if !ok {
		return 0, errors.New("cannot get ID")
	}

	ownerID := int(ownerIDstr.(float64))

	owner, err := m.DB.GetUserByID(ownerID)
	if err != nil {
		return 0, err
	}

	if owner.RoleID != 2 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "you are not owner"})
		return 0, err
	}

	return owner.ID, nil
}
