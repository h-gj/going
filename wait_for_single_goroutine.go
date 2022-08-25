package main

import (
	"fmt"
	"time"
)

func timeConsumingTask(finished chan bool) {
	fmt.Println("Running time consuming task...")
	time.Sleep(time.Second * 1)
	finished <- true
	fmt.Println("Finished time consuming task.")
}

func main() {
	finished := make(chan bool)
	go timeConsumingTask(finished)
	<-finished
}
