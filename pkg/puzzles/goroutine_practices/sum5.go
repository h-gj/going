package goroutine_practices

import "fmt"

func sum(arr []int, ch chan int) {
	total := 0
	for _, v := range arr {
		total += v
	}
	ch <- total
}

func Sum5() {
	chan1 := make(chan int, 4)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 3, 4, 5, 5, 1}
	// 0 3 6
	for i := 0; i < len(arr); i = i + 3 {
		go sum(arr[i:i+3], chan1)
	}

	output := make([]int, 4)
	for i := 0; i < 4; i++ {
		output[i] = <-chan1
	}
	close(chan1)

	fmt.Println(output)
}
