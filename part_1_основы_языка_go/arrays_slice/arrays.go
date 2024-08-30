package main

import (
	"errors"
	"fmt"
)

func main() {
	message := [3]string{"1", "2", "3"}
	fmt.Println("array messages:", message)
	fmt.Println("index 1 = 5:", message)
	_ = printMessage(message)
	fmt.Println(message)
}

func printMessage(messages [3]string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}

	fmt.Println(messages)
	messages[1] = "5"

	return nil
}
