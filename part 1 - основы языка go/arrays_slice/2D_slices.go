package main

import "fmt"

func main() {

	matrix1 := make([][]int, 10)

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			matrix1[y] = make([]int, 10)
			matrix1[y][x] = x
		}
		fmt.Println(matrix1[x])
	}
	fmt.Println("slice 1 2d matrix: ", matrix1)
	fmt.Println("********************")

	counter := 0
	for {
		counter++
		fmt.Println(counter)
		if counter >= 10000 {
			fmt.Println("before break value:", counter)
			break
		}
	}
}
