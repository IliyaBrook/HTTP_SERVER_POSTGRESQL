package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// примеры
	//ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	// Отменяем контекст, и гарантируем что функция отмены будет вызвана, как только main завершит выполнение, что бы предотвратить утечку ресурсов, связанных с контекстом.
	//defer cancel()
	//func() {
	//	time.Sleep(time.Millisecond * 100)
	//}()

	ctx = context.WithValue(ctx, "id", 1)

	parse(ctx)
}

func parse(ctx context.Context) {
	id := ctx.Value("id")
	fmt.Printf("read id from context: %d\n", id)

	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceeded")
			return
		}
	}
}
