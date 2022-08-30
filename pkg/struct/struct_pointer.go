package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	p := Point{3, 4}
	pp := &p // create a point in normal way (with a prepending & symbol)
	pp.x = 1
	(*pp).y = 2
	fmt.Printf("Point[x=%d, y=%d]\n", p.x, p.y)

	// create a pointer with new
	p1 := new(Point)
	p1.x = 10
	p1.y = 100
	fmt.Printf("Point[x=%d, y=%d]\n", p1.x, p1.y)
}
