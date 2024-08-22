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

	users := generateUsers(1000)

	wg := &sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go func() {
			err, errorMessage := saveUserInfo(user)
			if err != nil {
				fmt.Printf(errorMessage, err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Время выполнения: %v\n", duration)
}

func saveUserInfo(user sharable.User) (error, string) {
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
	return nil, ""
}

func generateUsers(count int) []sharable.User {
	users := make([]sharable.User, count)
	for i := 0; i < count; i++ {
		randActionIndex := rand.Intn(len(sharable.Actions))
		users[i] = sharable.User{
			Id:    i + 1,
			Email: fmt.Sprintf("user%d@ninja.go", i+1),
			Logs:  sharable.GenerateLogs(randActionIndex),
		}
	}
	return users
}

//func generateLogs(count int) []sharable.LogItem {
//	logs := make([]sharable.LogItem, count)
//	for i := 0; i < count; i++ {
//		logs[i] = sharable.LogItem{
//			Timestamp: time.Now(),
//			Action:    sharable.Actions[rand.Intn(len(sharable.Actions)-1)],
//		}
//	}
//	return logs
//}
