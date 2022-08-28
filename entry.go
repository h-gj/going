package main

import (
	"fmt"
	"going/puzzles"
	//"time"
	//"time"
)

type People struct {
	Name string
}

func (p People) Strings() string {
	return fmt.Sprintf("%v", p)
}

func main() {
	//ch := make(chan int, 1000)
	//done := make(chan bool)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//
	//}()
	//go func() {
	//	for {
	//		a, ok := <-ch
	//		if !ok {
	//			fmt.Println("close")
	//			done <- true
	//			return
	//		}
	//		fmt.Println("a: ", a)
	//	}
	//}()
	//fmt.Println("ok")
	////time.Sleep(time.Second * 100)
	//<- done

	//chan1 := make(chan bool, 1)
	//chan1 <- true
	//
	//fmt.Println(<-chan1)
	//fmt.Println(<-chan1)
	puzzles.Fibonacci(10)
}
