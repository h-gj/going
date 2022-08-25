package main

import (
	"fmt"
	. "going/pkg"
)

func main() {
	p := Test{Name: "gorgine", Age: 24}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
}
