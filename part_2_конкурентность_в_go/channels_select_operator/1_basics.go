package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// запись в канал может быть только когда канал инициализирован
	var msg chan string
	fmt.Println("uninitialized channel:", msg) // nil: uninitialized always equal to nil

	msg = make(chan string)
	fmt.Println("initialized channel:", msg) // 0xc000106060: memory address

	go func() {
		time.Sleep(time.Second * 2)
		msg <- "Канал ниндзя"
	}()
	//value := <-msg
	//log.Println(value)
	// or
	log.Println(<-msg)
}
