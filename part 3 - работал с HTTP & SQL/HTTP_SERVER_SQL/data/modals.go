package data

type DB struct {
	Users  []User  `json:"users"`
	Orders []Order `json:"orders"`
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Orders []int  `json:"orders"`
}

type Order struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
