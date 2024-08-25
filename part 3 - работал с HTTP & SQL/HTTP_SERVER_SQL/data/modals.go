package data

import "time"

type DbStruct struct {
	Users  []UserStruct  `json:"users"`
	Orders []OrderStruct `json:"orders"`
}

type UserStruct struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registered_at"`
}

type OrderStruct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	UserID      int     `json:"user_id"`
}
