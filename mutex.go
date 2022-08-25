package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	counter map[string]int
	lock    sync.Mutex
}

func (c *SafeCounter) incr(key string) {
	c.lock.Lock()
	c.counter[key]++
	c.lock.Unlock()
}

func (c *SafeCounter) get(key string) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.counter[key]
}

func main() {
	counter := SafeCounter{counter: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go counter.incr("a")
	}
	time.Sleep(time.Second)
	fmt.Println(counter.get("a"))
}
