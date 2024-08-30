package main

import (
	"log"
)

func main() {

	msg := make(chan string, 3)

	msg <- "Канал ниндзя"
	msg <- "Канал ниндзя 2"
	msg <- "Канал ниндзя 3"

	close(msg)
	var count int

	// стандартный синтаксис, когда у нас в ok возвращается что канал закрыт мы выходим из цикла с помощью break
	//for {
	//	value, ok := <-msg
	//	if !ok {
	//		log.Printf("Channel Num: %v is closed.", count)
	//		break
	//	} else {
	//		count++
	//	}
	//	fmt.Println(value)
	//}

	// упрощенный синтаксис

	for value := range msg {
		count++
		log.Printf("current chanel %s.", value)
	}
}
