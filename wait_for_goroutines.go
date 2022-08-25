package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func timeConsumingTask1() {
	defer wg.Done()
	fmt.Println("Running time consuming task...")
	time.Sleep(time.Second * 1)
	fmt.Println("Finished time consuming task.")
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go timeConsumingTask1()
	}
	wg.Wait()
}
