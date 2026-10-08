package main

import (
	"fmt"
	"reflect"
)

/*

A type must match the method name, capitalization, parameters, and return types exactly.

Capitalization also controls visibility across packages:

- Describe() starts with uppercase, so it is exported and can be used from other packages.
- describe() starts with lowercase, so it is unexported and available only inside its own package.

The interface type name follows the same rule:

type Describer interface { ... } // exported
type describer interface { ... } // package-private

*/

// Describer is an interface. Any type with a Describe() string method can be
// stored in a Describer variable.
type Describer interface {
	Describe() string
}

type Book struct {
	Title string
}

func (b Book) Describe() string {
	return "book: " + b.Title
}

type Dog struct {
	Name string
}

// This method has a pointer receiver, so *Dog implements Describer.
// Checking d == nil lets this method safely handle a nil *Dog.
func (d *Dog) Describe() string {
	if d == nil {
		return "no dog"
	}
	return "dog: " + d.Name
}

// Note:
// - Methods with a value receiver, func (d Dog) ..., belong to both Dog and *Dog.
// - Methods with a pointer receiver, func (d *Dog) ..., belong only to *Dog.
// So only *Dog implements Describer, not Dog

func main() {
	// An interface value conceptually contains two pieces of information:
	//
	//   dynamic type  — the concrete TYPE currently stored in the interface
	//   dynamic value — the concrete VALUE currently stored in the interface
	//
	// An interface is equal to nil only when BOTH pieces are nil.

	var nothing Describer
	// nothing has: dynamic type = nil, dynamic value = nil
	fmt.Println("nothing == nil:", nothing == nil)
	fmt.Printf("nothing's dynamic type: %T\n", nothing)
	fmt.Println("nothing's dynamic value: ", reflect.ValueOf(nothing))

	fmt.Println("\n\n")

	var item Describer = Book{Title: "Pride and Prejudice"}
	// item has: dynamic type = Book, dynamic value = Book{...}
	fmt.Println("item == nil:", item == nil)
	fmt.Printf("item's dynamic type: %T\n", item)
	fmt.Println(item.Describe())

	fmt.Println("\n\n")

	var dog *Dog = nil
	// dog itself is a nil pointer.
	fmt.Println("dog == nil:", dog == nil)

	fmt.Println("\n\n")

	var typedNil Describer = dog
	// typedNil has: dynamic type = *Dog, dynamic value = nil
	// Since the dynamic type exists, the interface itself is NOT nil.
	fmt.Println("typedNil == nil:", typedNil == nil)
	fmt.Printf("typedNil's dynamic type: %T\n", typedNil)
	fmt.Println("typedNil's dynamic value: ", reflect.ValueOf(typedNil))
	fmt.Println(typedNil.Describe())
}
