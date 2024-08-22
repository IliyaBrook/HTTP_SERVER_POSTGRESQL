package main

import (
	"fmt"
)

// этот способ используется в 90% случаях

type Age int

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

// printUserInfo method of User value receiver
func (u User) printUserInfo() {
	fmt.Println("User name:", u.name)
	fmt.Println("User age:", u.age)
	fmt.Println("User sex:", u.sex)
	fmt.Println("User weight:", u.weight)
	fmt.Println("User height:", u.height)
}

type isAdultReturn struct {
	message string
	isAdult bool
}

func (a Age) isAdult() isAdultReturn {
	if a != 0 {
		if a >= 18 {
			fmt.Println()
			//return "User is an adult", true
			return isAdultReturn{"User is an adult", true}
		}
	}

	//return "User is not an adult, or wrong age entries", false
	return isAdultReturn{"User is not an adult, or wrong age entries", false}
}

// printUserInfo method of User pointer receiver
//func (u *User) printUserInfo() {
//	fmt.Println("User name:", u.name)
//	fmt.Println("User age:", u.age)
//	fmt.Println("User sex:", u.sex)
//	fmt.Println("User weight:", u.weight)
//	fmt.Println("User height:", u.height)
//}

// Dumb database example with constructor

//type DumbDatabase struct {
//	m map[string]string
//}
//
//func NewDumbDatabase() *DumbDatabase {
//	return &DumbDatabase{
//		m: make(map[string]string),
//	}
//}

func NewUser(name string, age int, sex string, weight int, height int) User {
	return User{
		name:   name,
		age:    Age(age),
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func main() {
	// initialize user struct
	// эту структуру переиспользовать будет нельзя, так как она инициализирована [редко используется данный способ]
	user := struct {
		name   string
		age    int
		sex    string
		weight int
		height int
	}{
		"Vasya", 23, "Male", 75, 185,
	}
	fmt.Println("user structure: ", user)
	// formatted log
	fmt.Printf("%+v\n", user)
	// struct from user
	user2 := User{
		"Iliya Brook",
		36,
		"Male",
		64,
		174,
	}
	fmt.Printf("%+v\n", user2)
	fmt.Println("user 2 name:", user2.name)
	user3 := NewUser("Elena", 33, "Male", 55, 165)
	fmt.Println("log elena:", user3)
	fmt.Println("*****************************")
	fmt.Println("user 1 print")
	fmt.Println("*****************************")

	// example user creation
	user5 := NewUser("Alex", 29, "Male", 70, 180)
	fmt.Println("user 5 Alex:", user5)
	user5.printUserInfo()

	fmt.Println("*****************************")
	fmt.Println("user 2 print Dmitry")

	// creating a new user using NewUser function
	user4 := NewUser("Dmitry", 40, "Male", 78, 175)
	fmt.Println("user 4:", user4)
	user4.printUserInfo()

	fmt.Println("user 4 age:", user4.age)
	user4ResultAge := user4.age.isAdult()
	fmt.Println("user4 result message:", user4ResultAge.message)
	fmt.Println("*****************************")

}
