package db

import "time"

type DatabaseStruct struct {
	Users  []UserStruct    `json:"users" db:"users"`
	Orders []ProductStruct `json:"orders" db:"orders"`
}

type UserStruct struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	RegisteredAt time.Time `json:"registered_at" db:"registered_at"`
}

type ProductStruct struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
}

type UserOrdersStruct struct {
	UserId  int `json:"user_id" db:"user_id"`
	OrderId int `json:"order_id" db:"order_id"`
}
