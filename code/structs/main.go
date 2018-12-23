package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	leo := person{"Leo", "Vujanic"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	var steve person
	var john person
	john.firstName = "John"
	john.lastName = "Doe"

	fmt.Println(leo)
	fmt.Println(alex)
	// zero values example
	fmt.Printf("%+v\n", steve)

	fmt.Printf("%+v\n", john)
}
