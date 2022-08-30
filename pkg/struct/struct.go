package main

import "fmt"

//  Refer to => https://zetcode.com/golang/struct/

type Person struct {
	Name string
	Age  int
}

type Order struct {
	Id     int
	Person Person
}

type OrderForPromotedFields struct {
	Id int
	Person
}

func main() {
	// struct a struct with values
	p := Person{"Gorgine", 23}
	fmt.Printf("Hi, guys. I am %s and I am %d years old.\n", p.Name, p.Age)

	// struct a struct with field and value pairs
	p2 := Person{Age: 20, Name: "xxx"}
	fmt.Printf("Hi, guys. I am %s and I am %d years old.\n", p2.Name, p2.Age)

	// declare a struct and then assign values on fields
	var p3 Person
	p3.Name = "p3"
	p3.Age = 10
	fmt.Printf("Hi, guys. I am %s and I am %d years old.\n", p3.Name, p3.Age)

	// anonymous struct
	u := struct {
		A string
	}{
		A: "blabla",
	}
	fmt.Printf("This is a anony struct, value of field A is %s.\n", u.A)

	// nested struct
	o := Order{
		Id: 100,
		Person: Person{
			Age:  18,
			Name: "Gorilia",
		},
	}
	fmt.Printf("%s paid the order[%d].\n", o.Person.Name, o.Id)

	// struct with promoted fields
	o1 := OrderForPromotedFields{
		Id: 111,
		Person: Person{
			"namehewhat",
			12,
		},
	}
	fmt.Printf("%s paid the order[%d].\n", o1.Name, o1.Id)

	// struct functions fields..

}
