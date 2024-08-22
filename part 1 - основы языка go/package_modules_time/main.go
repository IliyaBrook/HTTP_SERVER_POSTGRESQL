package main

import (
	"context"
	"fmt"
	"github.com/zhashkevych/scheduler"
	"os"
	"os/signal"
	"time"
)

//////

//func main() {
//	square := shape.NewSquare(5)
//	//circle := shape.Circle{8}
//
//	shape.PrintShapeArea(square)
//	//shape.PrintShapeArea(circle)
//	shape.PrintInterface(square)
//	//shape.PrintInterface(circle)
//	/// print string
//	shape.PrintInterface("it is a string")
//	/// print number
//	shape.PrintInterface(500)
//}

func main() {
	ctx := context.Background()

	/// *** course code ///

	t := time.Now()
	fmt.Printf("Current time: %s\n", t.String())

	/// *** course code ///

	worker := scheduler.NewScheduler()
	worker.Add(ctx, parseSubscriptionData, time.Second*5)
	worker.Add(ctx, sendStatistics, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}

func parseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("subscription parsed successfuly at %s\n", time.Now().String())
}

func sendStatistics(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("statistics sent at %s\n", time.Now().String())
}
