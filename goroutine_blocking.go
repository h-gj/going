package main

import (
	"fmt"
	"time"
)

//func main() {
//	done := make(chan bool)
//	done <- true // deadlock
//	fmt.Println(<-done)
//}

//func main() {
//	done := make(chan bool)
//	go func() {
//		time.Sleep(time.Second)
//		done <- true
//	}()
//	fmt.Println(<-done)
//}

func main() {
	done := make(chan bool, 1)
	func() {
		time.Sleep(time.Second)
		done <- true
	}()
	fmt.Println(<-done)
}
