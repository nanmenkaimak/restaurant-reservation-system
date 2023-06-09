package repository

import (
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"time"
)

type DatabaseRepo interface {
	InsertUser(newUser models.Users) (int, error)
	GetAllUser() ([]models.Users, error)
	GetUserByID(id int) (models.Users, error)
	GetUsersByRole(role int) ([]models.Users, error)
	UpdateUser(id int, updatedUser models.Users) error
	DeleteUser(id int) error

	GetIDOfRoleByName(name string) (models.Roles, error)
	GetIDOfRestsTypeByName(name string) (models.TypeOfRestaurant, error)
	GetIDOfFoodCategoryByName(name string) (models.CategoryOfFood, error)

	InsertRestaurant(newRestaurant models.Restaurants) error
	GetAllRests() ([]models.Restaurants, error)
	GetRestByID(id int) (models.Restaurants, error)
	GetRestsByType(typeID int) ([]models.Restaurants, error)
	GetRestsByCity(city string) ([]models.Restaurants, error)
	GetRestsByAddress(address string) ([]models.Restaurants, error)
	GetRestsByAverageCheck(check int) ([]models.Restaurants, error)
	GetRestsByName(name string) ([]models.Restaurants, error)
	GetAllRestsByOwner(ownerID int) ([]models.Restaurants, error)

	InsertReservation(newReservation models.Reservations) error
	GetReservationsByCostumerID(costumerID int) ([]models.Reservations, error)
	GetReservationsByTableID(tableID int) ([]models.Reservations, error)
	GetReservationsByDate(comingTime time.Time) ([]models.Reservations, error)
	GetReservationsByNumOfGuests(numGuests int) ([]models.Reservations, error)

	InsertFood(newFood models.Food) error
	GetFoodsByRestID(restID int) ([]models.Food, error)
	GetFoodsByCategoryID(categoryID int) ([]models.Food, error)
	GetFoodsByName(name string) ([]models.Food, error)

	InsertTable(newUser models.Seats) error
	GetAllTableOfRest(restID int) ([]models.Seats, error)
	GetTableByID(id int) (models.Seats, error)

	Authenticate(email string, password string) (int, string, error)
}
