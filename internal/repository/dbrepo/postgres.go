package dbrepo

import (
	"errors"
	"fmt"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (m *postgresDBRepo) InsertUser(newUser models.Users) (int, error) {
	err := m.checkSimilarEmail(newUser.Email)
	if err != nil {
		return 0, err
	}
	password, err := hashPassword(newUser.Password)
	if err != nil {
		return 0, err
	}
	newUser.Password = password
	result := m.DB.Create(&newUser)
	return newUser.ID, result.Error
}

func (m *postgresDBRepo) GetAllUser() ([]models.Users, error) {
	var users []models.Users
	result := m.DB.Find(&users)
	return users, result.Error
}

func (m *postgresDBRepo) GetUserByID(id int) (models.Users, error) {
	var user models.Users
	result := m.DB.Where("id = ?", id).Find(&user)

	return user, result.Error
}

func (m *postgresDBRepo) GetUsersByRole(role int) ([]models.Users, error) {
	var user []models.Users
	result := m.DB.Where("role_id = ?", role).Find(&user)

	return user, result.Error
}

func (m *postgresDBRepo) UpdateUser(id int, updatedUser models.Users) error {
	var newUser models.Users
	password, err := hashPassword(updatedUser.Password)
	if err != nil {
		return err
	}
	updatedUser.Password = password
	result := m.DB.Model(&newUser).Where("id = ?", id).Updates(updatedUser)

	return result.Error
}

func (m *postgresDBRepo) DeleteUser(id int) error {
	var user models.Users
	result := m.DB.Delete(&user, id)

	return result.Error
}

func (m *postgresDBRepo) GetIDOfRoleByName(name string) (models.Roles, error) {
	var role models.Roles
	result := m.DB.Where("name = ?", name).Find(&role)

	return role, result.Error
}

func (m *postgresDBRepo) GetIDOfRestsTypeByName(name string) (models.TypeOfRestaurant, error) {
	var typeOfRest models.TypeOfRestaurant
	result := m.DB.Where("name = ?", name).Find(&typeOfRest)

	return typeOfRest, result.Error
}

func (m *postgresDBRepo) GetIDOfFoodCategoryByName(name string) (models.CategoryOfFood, error) {
	var foodCategory models.CategoryOfFood
	result := m.DB.Where("name = ?", name).Find(&foodCategory)

	return foodCategory, result.Error
}

func (m *postgresDBRepo) InsertRestaurant(newRestaurant models.Restaurants) error {
	result := m.DB.Create(&newRestaurant)
	return result.Error
}

func (m *postgresDBRepo) GetAllRests() ([]models.Restaurants, error) {
	var restaurants []models.Restaurants
	result := m.DB.Find(&restaurants)
	return restaurants, result.Error
}

func (m *postgresDBRepo) GetRestByID(id int) (models.Restaurants, error) {
	var restaurant models.Restaurants
	result := m.DB.Where("id = ?", id).Find(&restaurant)
	return restaurant, result.Error
}

func (m *postgresDBRepo) GetRestsByType(typeID int) ([]models.Restaurants, error) {
	var restaurants []models.Restaurants
	result := m.DB.Where("type_id = ?", typeID).Find(&restaurants)
	return restaurants, result.Error
}

func (m *postgresDBRepo) GetRestsByCity(city string) ([]models.Restaurants, error) {
	var restaurants []models.Restaurants
	result := m.DB.Where("city = ?", city).Find(&restaurants)
	if len(restaurants) == 0 {
		return nil, errors.New("in this city there are no restaurants")
	}
	return restaurants, result.Error
}

func (m *postgresDBRepo) GetRestsByAddress(address string) ([]models.Restaurants, error) {
	var restaurants []models.Restaurants
	result := m.DB.Where("address = ?", address).Find(&restaurants)
	return restaurants, result.Error
}

func (m *postgresDBRepo) GetRestsByAverageCheck(check int) ([]models.Restaurants, error) {
	var restaurants []models.Restaurants
	result := m.DB.Where("average_check <= ?", check).Find(&restaurants)
	return restaurants, result.Error
}

func (m *postgresDBRepo) GetRestsByName(name string) ([]models.Restaurants, error) {
	var restaurants []models.Restaurants
	result := m.DB.Where("name = ?", name).Find(&restaurants)
	return restaurants, result.Error
}

func (m *postgresDBRepo) GetAllRestsByOwner(ownerID int) ([]models.Restaurants, error) {
	var restaurants []models.Restaurants
	result := m.DB.Where("owner_id = ?", ownerID).Find(&restaurants)
	return restaurants, result.Error
}

func (m *postgresDBRepo) InsertReservation(newReservation models.Reservations) error {
	var cnt int64
	var reservation models.Reservations
	hoursBefore := newReservation.ComingTime.Add(-10 * time.Hour)
	hoursAfter := newReservation.ComingTime.Add(10 * time.Hour)
	m.DB.Model(&reservation).Where("coming_time >= ? and coming_time <= ? and seat_id = ?", hoursBefore, hoursAfter, newReservation.SeatID).Count(&cnt)
	if cnt > 0 {
		return errors.New("the table is already reserved")
	}

	seat, _ := m.GetTableByID(newReservation.SeatID)
	if seat.Capacity < newReservation.NumOfGuests {
		return errors.New(fmt.Sprintf("table cannot serve more guests than %d", seat.Capacity))
	}
	result := m.DB.Create(&newReservation)
	return result.Error
}

func (m *postgresDBRepo) GetReservationsByCostumerID(costumerID int) ([]models.Reservations, error) {
	var reservations []models.Reservations
	result := m.DB.Where("costumer_id = ?", costumerID).Find(&reservations)
	return reservations, result.Error
}

func (m *postgresDBRepo) GetReservationsByTableID(tableID int) ([]models.Reservations, error) {
	var reservations []models.Reservations
	result := m.DB.Where("seat_id = ?", tableID).Find(&reservations)
	return reservations, result.Error
}

func (m *postgresDBRepo) GetReservationsByDate(comingTime time.Time) ([]models.Reservations, error) {
	var reservations []models.Reservations
	result := m.DB.Where("coming_time = ?", comingTime).Find(&reservations)
	return reservations, result.Error
}

func (m *postgresDBRepo) GetReservationsByNumOfGuests(numGuests int) ([]models.Reservations, error) {
	var reservations []models.Reservations
	result := m.DB.Where("num_of_guests <= ?", numGuests).Find(&reservations)
	return reservations, result.Error
}

func (m *postgresDBRepo) InsertFood(newFood models.Food) error {
	result := m.DB.Create(&newFood)
	return result.Error
}

func (m *postgresDBRepo) GetFoodsByRestID(restID int) ([]models.Food, error) {
	var foods []models.Food
	result := m.DB.Where("restaurant_id = ?", restID).Find(&foods)
	return foods, result.Error
}

func (m *postgresDBRepo) GetFoodsByCategoryID(categoryID int) ([]models.Food, error) {
	var foods []models.Food
	result := m.DB.Where("category_id = ?", categoryID).Find(&foods)
	return foods, result.Error
}

func (m *postgresDBRepo) GetFoodsByName(name string) ([]models.Food, error) {
	var foods []models.Food
	result := m.DB.Where("name = ?", name).Find(&foods)
	return foods, result.Error
}

func (m *postgresDBRepo) InsertTable(newTable models.Seats) error {
	result := m.DB.Create(&newTable)
	return result.Error
}

func (m *postgresDBRepo) GetAllTableOfRest(restID int) ([]models.Seats, error) {
	var tables []models.Seats
	result := m.DB.Where("restaurant_id = ?", restID).Find(&tables)
	return tables, result.Error
}

func (m *postgresDBRepo) GetTableByID(id int) (models.Seats, error) {
	var table models.Seats
	result := m.DB.Where("id = ?", id).Find(&table)

	return table, result.Error
}

func (m *postgresDBRepo) Authenticate(email string, password string) (int, string, error) {
	var user models.Users

	m.DB.First(&user, "email = ?", email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}
	return user.ID, user.Password, nil
}

func (m *postgresDBRepo) checkSimilarEmail(email string) error {
	var cnt int64
	var users []models.Users
	m.DB.Model(&users).Where("email = ?", email).Count(&cnt)
	if cnt > 0 {
		return errors.New("user exists with this email")
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
