package main

import "fmt"

type FuncType func(name string) string

func (fc FuncType) say(words string) {
	fmt.Println(fc(words))
}

func english(words string) string {
	return fmt.Sprintf("I'll speak %s in English.", words)
}

func french(words string) string {
	return fmt.Sprintf("I'll speak %s in French.", words)
}

func main() {
	fc := FuncType(english)
	fc.say("hi")
	fc1 := FuncType(french)
	fc1.say("hello")
}
