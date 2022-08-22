package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "1111"
	messages <- "2222"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
