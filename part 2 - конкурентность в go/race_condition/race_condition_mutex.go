package main

import (
	"fmt"
	"sync"
)

type counter struct {
	counter int
	mu      *sync.Mutex
}

func (c *counter) inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *counter) value() int {
	c.mu.Lock()
	defer c.mu.Lock()
	return c.counter
}

func main() {
	c := counter{
		mu: new(sync.Mutex),
	}
	for i := 0; i < 1000; i++ {
		go func() {
			c.inc()
		}()
	}
	fmt.Println(c.value())
}
