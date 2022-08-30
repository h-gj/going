package puzzles

import (
	"fmt"
	"sync"
)

func PrintInTurn() {
	printLetter := make(chan bool)
	printNumber := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		i := 1
		for {
			select {
			case <-printNumber:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				printLetter <- true
				//printNumber <- false
			}

		}
	}()
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			if i > 'Z' {
				wg.Done()
				return
			}

			select {
			case <-printLetter:
				fmt.Print(string(i))
				i++
				printNumber <- true
			}
		}
	}(&wg)
	printNumber <- true
	wg.Wait()
	fmt.Println()
	//time.Sleep(time.Second)
}
