package main

import "fmt"

func main() {
	fmt.Println("start")
	fmt.Println(findMin(9, 5, 6, 1))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	minVal := numbers[0]

	for _, i := range numbers {
		if i < minVal {
			minVal = i
		}
	}

	return minVal
}
