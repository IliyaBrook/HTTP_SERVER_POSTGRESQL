package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	resBool, err := enterTheClub(61)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(resBool)
}

func enterTheClub(age int) (bool, error) {
	tooOld := "sorry man!, your too old"
	tooYoung := "sorry, your too young, maybe next time"
	if age >= 18 && age < 60 {
		return true, nil
	} else if age > 60 {
		return false, errors.New(tooOld)
	}
	return false, errors.New(tooYoung)
}
