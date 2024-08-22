package main

import "fmt"

// like objects in js
func main() {
	users := map[string]int{
		"Vasya":  15,
		"Petya":  23,
		"Kostya": 48,
	}

	for key, value := range users {
		fmt.Println("key:", key, "value:", value)
	}
	// rm keys Petya
	delete(users, "Petya")
	// append Sergey
	users["Sergey"] = 40
	// check keys
	value, isExist := users["Vasya"]
	fmt.Println("value:", value)
	fmt.Println("isExist:", isExist)
	fmt.Println("users:", users)
	// we can use internal func len
	fmt.Println("length users:", len(users))
}
