package main

import "fmt"

func main() {
	// unanimous func
	func() {
		fmt.Println("unanimous func")
	}()
	incF := increment()
	varInc := incF()
	varInc = incF()
	varInc = incF()
	fmt.Println(varInc)
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
