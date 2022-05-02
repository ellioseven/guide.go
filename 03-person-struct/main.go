// Defines package as an executable.
package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jim@example",
			zip:   1000,
		},
	}

	// go passes copies as function arguments by default for some types, if the
	// function implements a reference as a type, it will be passed as a reference
	// instead of a copied value, which can be used to update the value in memory
	// in place, note: some types are exclusive from this rule: slices, maps,
	// channels, pointers and functions.
	jim.setName("Jake")

	jim.print()
}

func (_person person) print() {
	fmt.Printf("%+v", _person)
}

// `*person` is a pointer type declaration, `*_person` is the resolved value
// of the pointer, which is used manipulate the value of the reference, turn
// adress into value with `*address`, turn value into address with `&value`
func (_person *person) setName(firstName string) {
	(*_person).firstName = firstName
}
