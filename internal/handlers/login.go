package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/JWT"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"net/http"
)

func (m *Repository) Login(ctx *gin.Context) {
	var user models.Users

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ctx.Error(err))
		return
	}

	id, _, err := m.DB.Authenticate(user.Email, user.Password)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ctx.Error(err))
		return
	}

	token, err := JWT.GenerateToken(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ctx.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
