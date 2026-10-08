package main

import "fmt"

// Person groups related values into one custom type.
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a Person value and set its fields by name.
	person := Person{
		Name: "Avery",
		Age:  25,
	}

	// Read individual fields with dot notation.
	fmt.Println(person.Name)
	fmt.Println(person.Age)

	// Change one field on the struct value.
	person.Age++
	fmt.Printf("%s is now %d\n", person.Name, person.Age)
}
