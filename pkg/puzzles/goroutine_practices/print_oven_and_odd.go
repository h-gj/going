package goroutine_practices

import "fmt"

func printOdd(ch chan int) {
	for v := range ch {
		fmt.Printf("Odd: %v\n", v)
	}
}

func printEven(ch chan int) {
	for v := range ch {
		fmt.Printf("Even: %v\n", v)
	}
}

func PrintOvenAndOdd() {
	num := []int{2, 1, 54, 543543, 4352342, 234234, 4324, 234324, 23434, 2323, 23, 5, 4}
	chOdd := make(chan int)
	chEven := make(chan int)

	go printOdd(chOdd)
	go printEven(chEven)

	for _, v := range num {
		if v%2 == 0 {
			chEven <- v
		} else {
			chOdd <- v
		}
	}

}
