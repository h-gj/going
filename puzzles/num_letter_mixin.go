package puzzles

import "fmt"

func PrintString(r rune) string {
	a := "fdsfdsfdsf"
	b := []rune(a)
	fmt.Println(b)
	c := []rune{'a', 'b'}
	fmt.Println(c)
	return string(a)

	//switch {
	//case 97 <= r && r <= 122:
	//	return r - 32
	//case 65 <= r && r <= 90:
	//	return r + 32
	//default:
	//	return r
	//}

	//lock := sync.WaitGroup{}
	//printNumber := make(chan bool)
	//printLetter := make(chan bool)
	//
	//go func() {
	//	i := 1
	//	for {
	//		select {
	//		case <-printNumber:
	//			fmt.Print(i)
	//			i++
	//			fmt.Print(i)
	//			i++
	//			printLetter <- true
	//		}
	//	}
	//}()
	//lock.Add(1)
	//
	//go func(lock *sync.WaitGroup) {
	//	i := 'A'
	//	for {
	//		select {
	//		case <-printLetter:
	//			if i > 'Z' {
	//				lock.Done()
	//				return
	//			}
	//			fmt.Print(string(i))
	//			i++
	//			printNumber <- true
	//		}
	//
	//	}
	//}(&lock)
	//printNumber <- true
	//lock.Wait()

}
