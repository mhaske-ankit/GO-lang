package main

import "fmt"

// Define a struct named Person
type Person struct {
	name string
	age  int
}

// Define a method named greet for Person
func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func main() {
	// Create a Person instance
	person := Person{
		name: "Alice",
		age:  30,
	}

	// Call the greet method of the Person instance
	person.greet()
}
