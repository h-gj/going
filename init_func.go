package main

import "fmt"

func init() {
	fmt.Println("running init...")
}

func main() {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}
