package main

import (
	"fmt"
	"time"
)

func chanOp(c chan string) {
	fmt.Println("chan op start")
	time.Sleep(time.Second)
	c <- "1"
	fmt.Println("chan op end")
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	//chanOp(chan1)
	select {
	case msg := <-chan1:
		fmt.Println(msg)
	default:
		fmt.Println("no msg")
	}

	msg := "msg for sent"
	select {
	case chan2 <- msg:
		fmt.Println("sent msg: ", msg)
	default:
		fmt.Println("no msg for sending")
	}

}
