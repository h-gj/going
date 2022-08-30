package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Start working...")
	time.Sleep(time.Second)
	fmt.Println("Finish working...")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	res := <-done
	fmt.Println(res)
}
