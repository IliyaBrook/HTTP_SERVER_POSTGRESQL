package main

import (
	"errors"
	"fmt"
)

func main() {
	// init slice
	var messages []string
	fmt.Println("log empty slice:", messages)
	messages = append(messages, "first index")
	fmt.Println("after add first index:", messages)

	messages = append(messages, "index 22 add")
	fmt.Println("after add index 2:", messages)
	fmt.Println("****** without init slice ******")
	// noInitSlice := []string{}
	var noInitSlice []string
	// recommended to use  	var noInitSlice []string, because this syntax initialize slice as nill and not use memory resource as []string{}
	fmt.Println(noInitSlice)

	fmt.Println("****** using make ******")

	_ = printMessage2(messages)

	// make, not init slice
	makeMessages := make([]string, 5)
	fmt.Println(makeMessages)
}

func printMessage2(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}

	messages[1] = "5"

	fmt.Println(messages)

	return nil
}
