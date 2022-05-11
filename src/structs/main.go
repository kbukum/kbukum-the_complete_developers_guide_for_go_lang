package main

import "fmt"

type person1 struct {
	firstName string
	lastName  string
}

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contact
}

func (p person) print() {
	fmt.Printf("Person = %+v\n", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func main() {

	fmt.Println("# Struct Examples")
	fmt.Println()
	fmt.Println("## Creating Struct")
	fmt.Println()
	alex := person1{"Alex", "Anderson"}
	fmt.Printf("Struct for Alex: %+v\n", alex)

	simon := person1{firstName: "Simon", lastName: "Celeb"}
	fmt.Printf("Struct for Simon: %+v\n", simon)

	var empty person1
	fmt.Printf("Struct for Empty: %+v\n", empty)
	fmt.Println()
	fmt.Println("## Embedding Struct")
	fmt.Println()
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contact{
			email: "jim@test.com",
			zip:   32131,
		},
	}
	fmt.Printf("Struct for Embedding Types: %+v\n", jim)
	fmt.Println()
	fmt.Println("## Adding receiver function to struct")
	jim.print()
	jim.updateName("Jimmy")
	fmt.Println("Jim name is updated")
	jim.print()
}
