package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Notifier stores a channel as one of its fields. The channel belongs to this
// particular Notifier value, just like its Name field does.
type Notifier struct {
	Name     string
	// this is a channel that can send and receive strings. It is owned by this Notifier.
	Messages chan string
}

// NewNotifier constructs a Notifier and initializes its channel field.
func NewNotifier(name string) *Notifier {
	return &Notifier{
		Name:     name,
		Messages: make(chan string, 5),
	}
}

const (
	min = 1000
	max = 10000
)

// Send puts a message into this notifier's channel.
func (n *Notifier) Send(message string) {
	fmt.Printf("%s sending: %s\n", n.Name, message)
	random := min + rand.Intn(max-min+1) // random number between min and max
	time.Sleep(time.Duration(random) * time.Millisecond) // simulate some work
	n.Messages <- message
}

// Listen receives and prints every message until Messages is closed.
func (n *Notifier) Listen(wg *sync.WaitGroup) {
	defer func() {
		fmt.Printf("%s stopping listening\n", n.Name)
		wg.Done()
	}()

	fmt.Printf("%s starting to listen\n", n.Name)

	for message := range n.Messages {
		fmt.Printf("%s received: %s\n", n.Name, message)
	}
}

func main() {
	// This struct now owns a channel named Messages.
	notifier := NewNotifier("console notifier")

	var listenerWG sync.WaitGroup
	var senderWG sync.WaitGroup
	
	listenerWG.Add(1)
	// start a goroutine that listens for messages on the struct's channel field.
	go notifier.Listen(&listenerWG)

	// send 5 messages all with goroutines
	for i := 1; i <= 5; i++ {
		senderWG.Go(func() {
			notifier.Send(fmt.Sprintf("message %d", i))
		})
	}

	// wait for all senders to finish
	senderWG.Wait() 
	// after all the sending is done, close the channel
	close(notifier.Messages)
	// wait for the listener to finish
	listenerWG.Wait()
}
