package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/forms"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
	"time"
)

func (m *Repository) SignUp(ctx *gin.Context) {
	var newUser models.Users

	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form := forms.New()

	form.IsEmail(newUser.Email)
	form.MinLength(newUser.Password, 8)

	if !form.Valid() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "your values is not valid"})
		return
	}

	_, err := m.DB.InsertUser(newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}
