package main

import "fmt"

// addOne receives a pointer to an int. *number means "the int stored at the
// address in number", so changing it changes the original variable.
func addOne(number *int) {
	*number = *number + 1
}

func main() {
	count := 10

	// &count means "the address of count". Store that address in a pointer.
	countPointer := &count

	fmt.Println("value:", count)
	fmt.Println("address:", countPointer)

	// *countPointer means "the value at this address" (dereferencing).
	*countPointer = 20
	fmt.Println("after dereferencing:", count)

	// Pass the address to a function so it can update the same original value.
	addOne(&count)
	fmt.Println("after addOne:", count)
}
