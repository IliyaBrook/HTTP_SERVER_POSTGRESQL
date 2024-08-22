package main

import (
	"Golang/sharable"
	"Golang/utils"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// without goroutines 15 s
// with goroutines 81.4ms
func main() {
	start := time.Now()

	users := make(chan sharable.User)
	go generateUsersChan(1000, users)

	wg := &sync.WaitGroup{}
	for user := range users {
		wg.Add(1)
		go func() {
			err, errorMessage := saveUserInfoChan(user, wg)
			if err != nil {
				fmt.Printf(errorMessage, err)
			}
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Время выполнения: %v\n", duration)
}

func saveUserInfoChan(user sharable.User, wg *sync.WaitGroup) (error, string) {
	// delay 10 ms to emulate network delays
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("WRITTING FILE FOR USER ID: %d\n", user.Id)
	currentDir := utils.GetFilePath()
	filename := fmt.Sprintf("%s/logs/uid_%d.txt", currentDir, user.Id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err, "failed to open file"
	}

	_, err = file.WriteString(user.GetActivityInfo())
	if err != nil {
		return err, "failed to write to file"

	}
	wg.Done()
	return nil, ""
}

func generateUsersChan(count int, users chan sharable.User) {
	for i := 0; i < count; i++ {
		randActionIndex := rand.Intn(len(sharable.Actions))
		users <- sharable.User{
			Id:    i + 1,
			Email: fmt.Sprintf("user%d@ninja.go", i+1),
			Logs:  sharable.GenerateLogs(randActionIndex),
		}
		//time.Sleep(time.Millisecond * 10)
	}
	close(users)
}
