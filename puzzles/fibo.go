package puzzles

import "fmt"

func FibonacciFirstN(ch chan int, ch1 chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-ch1:
			return
		}
	}
}

func Fibonacci(firstN int) {
	ch := make(chan int)
	ch1 := make(chan bool)
	go func() {
		for i := 0; i < firstN; i++ {
			fmt.Println(<-ch)
		}
		ch1 <- true
	}()
	FibonacciFirstN(ch, ch1)

}
