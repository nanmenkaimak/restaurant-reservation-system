package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func (m *Repository) checkOwner(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get("id")
	if !ok {
		return 0, errors.New("cannot get ID")
	}

	ownerID := int(id.(float64))

	owner, err := m.DB.GetUserByID(ownerID)
	if err != nil {
		return 0, errors.New("cannot get user from database")
	}

	if owner.RoleID != 2 {
		return 0, errors.New("only owner can add")
	}

	return ownerID, nil
}
