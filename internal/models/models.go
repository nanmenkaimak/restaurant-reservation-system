package models

import "time"

type Users struct {
	ID        int       `json:"id" gorm:"primaryKey; autoIncrement:1"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleID    int       `json:"role_id"`
	Role      Roles     `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Roles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Restaurants struct {
	ID           int              `json:"id" gorm:"primaryKey; autoIncrement:1"`
	Name         string           `json:"name"`
	OwnerID      int              `json:"owner_id"`
	Owner        Users            `json:"owner"`
	TypeID       int              `json:"type_id"`
	Type         TypeOfRestaurant `json:"type"`
	AverageCheck int              `json:"average_check"`
	City         string           `json:"city"`
	Address      string           `json:"address"`
	About        string           `json:"about"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

type Seats struct {
	ID           int         `json:"id" gorm:"primaryKey; autoIncrement:1"`
	RestaurantID int         `json:"restaurant_id"`
	Restaurant   Restaurants `json:"restaurant"`
	Capacity     int         `json:"capacity"`
}

type TypeOfRestaurant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Food struct {
	ID           int            `json:"id" gorm:"primaryKey; autoIncrement:1"`
	Name         string         `json:"name"`
	RestaurantID int            `json:"restaurant_id"`
	Restaurant   Restaurants    `json:"restaurant"`
	CategoryID   int            `json:"category_id"`
	Category     CategoryOfFood `json:"category"`
}

type CategoryOfFood struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Reservations struct {
	ID          int       `json:"id" gorm:"primaryKey; autoIncrement:1"`
	CostumerID  int       `json:"costumer_id"`
	Costumer    Users     `json:"costumer"`
	SeatID      int       `json:"seat_id"`
	Seat        Seats     `json:"seat"`
	ComingTime  time.Time `json:"coming_time"`
	NumOfGuests int       `json:"num_of_guests"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
