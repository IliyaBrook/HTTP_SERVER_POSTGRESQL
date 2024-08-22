package main

import "fmt"

func init() {
	fmt.Println("initialize func running before main func")
}

func main() {
	message := "I'm developer"
	changeMessage(&message)

	fmt.Println(message)
}

func changeMessage(message *string) {
	*message += " and im daddy"
	num := 100
	fmt.Println(message)
	fmt.Println("log memory num:", &num)
}
