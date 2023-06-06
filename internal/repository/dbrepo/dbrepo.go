package dbrepo

import (
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/repository"
	"gorm.io/gorm"
)

type postgresDBRepo struct {
	DB *gorm.DB
}

func NewPostgresRepo(conn *gorm.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}
