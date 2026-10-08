package main

import (
	"fmt"
	"sync"
	"time"
)

/*
chan T      // can send and receive values of T
<-chan T    // receive-only channel
chan<- T    // send-only channel
*/

// worker is run by each worker goroutine.
//
// jobs <-chan int means this function may only RECEIVE ints from jobs.
// results chan<- string means this function may only SEND strings to results.
// The arrows make the intended direction of data flow explicit.
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {

	fmt.Printf("worker %d starting\n", id)

	// Tell the WaitGroup that this specific worker has finished when the function returns.
	// this is an anonymouns func that is invoked immediately.
	defer func() {
		fmt.Printf("worker %d stopping\n", id)
		wg.Done()
	}()

	// Keep receiving jobs until the jobs channel is closed and all queued jobs
	// have been received.
	for job := range jobs {
		// Make a job take as many seconds as its number. For example, job 3
		// pauses this worker for three seconds. While it sleeps, the other
		// workers can receive and process other jobs.
		fmt.Printf("worker %d started job %d (%d seconds)\n", id, job, job)
		time.Sleep(time.Duration(job) * time.Second)

		// Send the completed work to the results channel.
		results <- fmt.Sprintf("worker %d finished job %d", id, job)
	}
}

func main() {
	// Create two unbuffered channels. A send waits until another goroutine is
	// ready to receive, which coordinates the goroutines without explicit locks.

	// buffered means there is a "waiting area" inside the channel
	// if there is no receiver ready, it will wait until there is one
	// its size is the second argument of make
	// unbuffered means there is no waiting area, so a receiver must be ready
	// but in this the sender runs in its own goroutine, so if it reaches jobs <- 1 before
	// a worker starts waiting, it simply pauses until a worker is ready to receive the job
	jobs := make(chan int)
	results := make(chan string)

	// A WaitGroup lets us wait until all workers have stopped.
	var wg sync.WaitGroup

	// Start THREE worker goroutines. They all receive from the same jobs channel,
	// so each job is handled by exactly one available worker.
	for id := 1; id <= 3; id++ {
		// tell the WaitGroup that we have one more worker to wait for
		wg.Add(1)
		// Start a worker goroutine. It will call wg.Done() when it finishes
		go worker(id, jobs, results, &wg)
	}

	// This goroutine supplies jobs, then shuts down the pipeline in order.
	go func() {
		// Send six integers into the jobs channel.
		for job := 1; job <= 6; job++ {
			jobs <- job
		}

		// Closing jobs tells workers there will be no more jobs. Their range loops
		// end after they finish any job they already received.
		close(jobs)

		// Wait until every worker has called wg.Done(). Only then is it safe to
		// close results, because no worker can try to send another result.
		wg.Wait()
		close(results)

	}()

	// Print results as they arrive. This range ends automatically once results
	// is closed and every previously sent result has been received.
	// The output order may vary because workers run concurrently.
	for result := range results {
		fmt.Println(result)
	}
}
