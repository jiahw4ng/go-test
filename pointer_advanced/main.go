package main

import "fmt"

// Node is one item in a linked list.
// Next holds the address of the next Node. A nil Next means this is the last
// node in the list.
type Node struct {
	Value string
	Next  *Node
}

// LinkedList stores pointers to the first and last node.
type LinkedList struct {
	Head *Node
	Tail *Node
}

// Append creates a new Node and links it after the current last node.
// *LinkedList is a pointer receiver because Append must update Head and Tail
// on the original list.
func (list *LinkedList) Append(value string) {
	newNode := &Node{Value: value} // &Node creates a Node and returns its address.

	if list.Head == nil {
		// The list is empty, so this node is both the first and last node.
		list.Head = newNode
		list.Tail = newNode
		return
	}

	// Tail points to the current final node. Change that node's Next pointer so
	// it points at newNode, then move Tail forward.
	list.Tail.Next = newNode
	list.Tail = newNode
}

// Find returns a pointer to the matching node. Returning a pointer allows the
// caller to change the real node inside the list, rather than a copied Node.
func (list *LinkedList) Find(value string) *Node {
	// current is a pointer that walks from node to node.
	for current := list.Head; current != nil; current = current.Next {
		if current.Value == value {
			return current
		}
	}

	return nil // no matching node exists
}

// Print follows pointers through the list until it reaches nil.
func (list *LinkedList) Print() {
	for current := list.Head; current != nil; current = current.Next {
		fmt.Printf("%s -> ", current.Value)
	}
	fmt.Println("nil")
}

func main() {
	var languages LinkedList // Head and Tail start as nil pointers.

	languages.Append("Go")
	languages.Append("Rust")
	languages.Append("Python")

	fmt.Print("Before: ")
	languages.Print()

	// Find returns a pointer to the actual "Rust" node. Updating through that
	// pointer changes the node stored in the list.
	if language := languages.Find("Rust"); language != nil {
		language.Value = "TypeScript"
	}

	fmt.Print("After:  ")
	languages.Print()
}
