package handlers

import (
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/repository"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/repository/dbrepo"
	"gorm.io/gorm"
)

// Repo is repository used by handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	DB repository.DatabaseRepo
}

// NewRepo creates new repository
func NewRepo(db *gorm.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db),
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
