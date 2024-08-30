package main

import "fmt"

func main() {
	//panic("panic")
	// defer func runs after the function ends
	//defer func() {
	//	fmt.Println("some message")
	//}()

	defer handlerPanic()
	// panic logic
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	messages[4] = "message 5"

}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
