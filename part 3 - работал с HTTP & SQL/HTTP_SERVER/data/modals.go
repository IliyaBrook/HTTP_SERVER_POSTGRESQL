package data

type DB struct {
	Users []User `json:"users"`
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}
