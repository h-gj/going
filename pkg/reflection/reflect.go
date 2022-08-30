package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (t *T) Hello(name string) {
	fmt.Printf("hello %s", name)
}

func main() {
	t := &T{}
	in := []reflect.Value{reflect.ValueOf("fdsf")}
	reflect.ValueOf(t).MethodByName("Hello").Call(in)
}
