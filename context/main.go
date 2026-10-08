package main

import (
	"context"
	"fmt"
	"time"
)

/*

context is the standard way to pass control information through go function calls and goroutines
cancellation: "stop now"
timeout: "stop after a certain time"

mental model: "this request is still alive" signal

does not forcibly kill a goroutine, instead use

case <-ctx.Done(): to check if the context has been cancelled or timed out
	return ctx.Err() to explain why the context ended, such as "context deadline exceeded"

*/

// doWork performs steps until it finishes or its context is cancelled.
func doWork(ctx context.Context) error {
	for step := 1; step <= 5; step++ {
		// select waits on multiple channel operations and runs the case
		// that becomes ready first.
		// - ctx.Done(): the timeout expired or somebody called cancel.
		// - time.After(...): this simulated step finishes after one second.
		select {
		case <-ctx.Done():
			// ctx.Err() explains why the context ended, such as
			// "context deadline exceeded".
			return ctx.Err()
		case <-time.After(time.Second):
			fmt.Println("finished step", step)
		}
	}

	fmt.Println("all steps completed")

	return nil
}

func main() {
	// Start from context.Background(), then derive a context that automatically
	// cancels after 2500 milliseconds. The work needs five seconds, so it will
	// be stopped partway through.

	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Millisecond)

	// Always call cancel when we are done. It releases context-related resources
	// early if doWork finishes before the timeout.
	defer cancel()

	if err := doWork(ctx); err != nil {
		fmt.Println("work stopped:", err)
		return
	}

	fmt.Println("work completed")
}
