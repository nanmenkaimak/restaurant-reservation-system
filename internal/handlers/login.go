package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/JWT"
	"net/http"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary Login
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body LoginParams true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /user/login [post]
func (m *Repository) Login(ctx *gin.Context) {
	var user LoginParams

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _, err := m.DB.Authenticate(user.Email, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := JWT.GenerateToken(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
