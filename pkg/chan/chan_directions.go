package main

import "fmt"

func acceptOnly(msg chan<- string, val string) {
	msg <- val
}

func sendOnly(chan1 <-chan string, chan2 chan<- string) {
	chan1Res := <-chan1
	chan2 <- chan1Res
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	acceptOnly(chan1, "22222")
	sendOnly(chan1, chan2)
	//chan2Res := <-chan2
	fmt.Println(<-chan2)
}
