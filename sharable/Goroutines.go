package sharable

import (
	"fmt"
	"math/rand"
	"time"
)

var Actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type LogItem struct {
	Action    string
	Timestamp time.Time
}

type User struct {
	Id    int
	Email string
	Logs  []LogItem
}

func (u User) GetActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.Id, u.Email)
	for i, item := range u.Logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.Action, item.Timestamp)
	}
	return out
}

func GenerateLogs(count int) []LogItem {
	logs := make([]LogItem, count)
	for i := 0; i < count; i++ {
		logs[i] = LogItem{
			Timestamp: time.Now(),
			Action:    Actions[rand.Intn(len(Actions)-1)],
		}
	}
	return logs
}
