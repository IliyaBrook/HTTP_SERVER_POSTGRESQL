package data

type DB struct {
	Users  []User  `json:"users"`
	Orders []Order `json:"orders"`
}

type User struct {
	ID     int    `json:"ID"`
	Name   string `json:"Name"`
	Orders []int  `json:"Orders"`
}

type Order struct {
	ID          int     `json:"ID"`
	UserID      int     `json:"UserID"`
	Name        string  `json:"Name"`
	Quantity    int     `json:"Quantity"`
	Price       float64 `json:"Price"`
	Description string  `json:"Description"`
}
