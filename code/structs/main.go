package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
}

type personWithContact struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// maps contactInfo prop to contactInfo type
// type personWithContact struct {
// 	firstName string
// 	lastName string
// 	contactInfo
// }

func main() {

	justPerson()

	withContact()

	changeUsername()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func justPerson() {
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

func withContact() {
	alex := personWithContact{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@anderson.mail.loc",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v\n", alex)
}

func changeUsername() {
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Printf("%+v\n", alex)

	// alexPointer := &alex
	// alexPointer.updateName("Pero")
	// call directly on type or can be called over pointer
	alex.updateName("Pero")

	fmt.Printf("%+v\n", alex)
}
