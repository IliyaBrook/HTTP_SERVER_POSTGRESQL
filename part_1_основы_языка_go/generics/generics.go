package main

import (
	"fmt"
	"reflect"
)

type Number interface {
	int64 | float64
}

type Number2 int64

func (n *Number2) Number2Incr() Number2 {
	*n += 1
	return *n
}

type User struct {
	Email string
	Name  string
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}
	int64Num2 := []Number2{1, 2, 3}

	users := []User{
		{
			"iliyabrook1987@gmail.com",
			"Iliya Brook",
		},
		{
			"nikabrook@gmail.com",
			"veronika Brook",
		},
	}

	//for _, user := range d {
	//	fmt.Printf("%+v\n", user)
	//}

	var accSum Number2
	for range int64Num2 {
		accSum.Number2Incr()
	}
	fmt.Println("acc sum num 2 res:", accSum)

	genInt64SumRes := sum(a)
	genFloat64SumRes := sum(b)
	fmt.Println("sum result int64:", genInt64SumRes)
	fmt.Println("sum result float64:", genFloat64SumRes)

	// generic comparable and any
	comparableResult := searchElement(c, "22")
	fmt.Println("comparable result: ", comparableResult)

	// find nika result
	fmt.Println("********* find Nika user result *********")
	fmt.Printf("%+v\n", searchElement(users,
		User{
			"nikabrook@gmail.com",
			"veronika Brook",
		},
	))

	fmt.Println("********* find Nika user result [Any] *********")
	fmt.Printf("%+v\n", searchElementAny(users,
		"nikabrook@gmail.com",
	))
}

// func sum[V cmp.Ordered](array []V) V { if we want all types

func sum[V Number](array []V) V {
	//accum := int64(0)
	var accum V
	for _, i := range array {
		accum += i
	}
	return accum
}

// generic comparable and any

func searchElement[C comparable](elements []C, searchE C) bool {

	for _, elem := range elements {
		if elem == searchE {
			return true
		}
	}
	return false
}

func searchElementAny[C comparable](elements []C, searchE any) bool {
	for _, elem := range elements {
		searchEType := reflect.TypeOf(searchE)
		if elem == searchE {
			return true
		} else if searchEType == reflect.TypeOf("") {
			elemValue := reflect.ValueOf(elem)
			if elemValue.Kind() == reflect.Struct {
				for index := range elemValue.NumField() {
					field := elemValue.Field(index)
					fieldStr, ok := field.Interface().(string)
					if ok && fieldStr == searchE {
						return true
					}
				}
			}
		}
	}
	return false
}
