package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GOLANG-NINJA/pingrobot/workerpool"
)

const (
	Interval       = time.Second * 10
	RequestTimeout = time.Second * 3
	WorkersCount   = 3
)

var urls = []string{
	"https://workshop.zhashkevych.com/",
	"https://golang-ninja.com/",
	"https://zhashkevych.com/",
	"https://google.com/",
	"https://golang.org/",
}

func main() {
	results := make(chan workerpool.Result)
	workerPool := workerpool.New(WorkersCount, RequestTimeout, results)

	workerPool.Init()

	go generateJobs(workerPool)
	go processResults(results)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	workerPool.Stop()
}

func processResults(results chan workerpool.Result) {
	go func() {
		for result := range results {
			fmt.Println(result.Info())
		}
	}()
}

func generateJobs(wp *workerpool.Pool) {
	for {
		for _, url := range urls {
			wp.Push(workerpool.Job{URL: url})
		}

		time.Sleep(Interval)
	}
}
