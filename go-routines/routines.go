package main

import (
	"fmt"
	"time"
)

/*
	GOROUTINES + CHANNELS CHEAT-SHEET

	1) Unbuffered channel: make(chan T)
	   - Send blocks until a receiver is ready (synchronous handoff).
	   - Great for "rendezvous" / coordination.

	2) Buffered channel: make(chan T, N)
	   - Send blocks only when buffer is full.
	   - Receive blocks only when buffer is empty.
	   - Useful for smoothing bursts / limited queue.

	3) select:
	   - Waits for one of multiple channel ops to become ready.
	   - If multiple are ready, one is chosen pseudo-randomly.
	   - You can add:
	     - default: non-blocking select
	     - time.After: timeout

	4) "done" / cancellation:
	   - A channel is used to signal "stop".
	   - Convention: close(done) broadcasts cancellation to all receivers.
	   - Prefer `chan struct{}` for done signals (no payload needed).

	5) Pipelines:
	   - Each stage is a goroutine:
	     input chan -> transform -> output chan
	   - Stages must close their output when finished (very important).
*/

// ------------------------------------------------------------
// 1) SELECT with TWO UNBUFFERED channels
// ------------------------------------------------------------

func selectTwoChannels() {
	myChannel := make(chan string)      // unbuffered: send waits for receiver
	anotherChannel := make(chan string) // unbuffered: send waits for receiver

	// `go func(){...}()` is:
	// - anonymous function definition: func(){...}
	// - immediate call: ()
	// - running in a goroutine: go
	go func() {
		myChannel <- "Data"
	}()

	go func() {
		anotherChannel <- "Another Data"
	}()

	// select receives from whichever channel becomes ready first.
	// NOTE: the other goroutine may remain blocked forever if we don't receive its value.
	select {
	case msg := <-myChannel:
		fmt.Println("Received from myChannel:", msg)
	case msg := <-anotherChannel:
		fmt.Println("Received from anotherChannel:", msg)
	}
}

// A safer version: drain both so no goroutine is left blocked.
func selectTwoChannelsSafe() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() { myChannel <- "Data" }()
	go func() { anotherChannel <- "Another Data" }()

	// Receive BOTH values (order is still whichever comes first)
	for i := 0; i < 2; i++ {
		select {
		case msg := <-myChannel:
			fmt.Println("Received from myChannel:", msg)
		case msg := <-anotherChannel:
			fmt.Println("Received from anotherChannel:", msg)
		}
	}
}

// ------------------------------------------------------------
// 2) BUFFERED CHANNEL (capacity N)
// ------------------------------------------------------------

func bufferedChannelExample() {
	// Buffered channel with capacity 3
	// - up to 3 sends can happen without a receiver
	charCh := make(chan string, 3)

	chars := []string{"A", "B", "C"}

	// Here select is unnecessary (only one case), but keeping it to match the tutorial style.
	for _, c := range chars {
		select {
		case charCh <- c:
			// Sent into buffer
		}
	}

	// Always close the channel when:
	// - you are the sender
	// - and you know no more values will be sent
	close(charCh)

	// range over channel reads until it is closed AND drained
	for v := range charCh {
		fmt.Println("Received:", v)
	}
}

// ------------------------------------------------------------
// 3) DONE / CANCELLATION CHANNEL
// ------------------------------------------------------------

// Convention: use `chan struct{}` for done signals (no payload, minimal alloc).
func doWork(done <-chan struct{}) {
	// This loop runs forever until done is closed.
	for {
		select {
		case <-done:
			// When `done` is closed, receive unblocks immediately.
			fmt.Println("Stopping work.")
			return

		default:
			// default makes the select non-blocking.
			// WARNING: this becomes a busy-loop (CPU-heavy) unless you sleep/yield.
			fmt.Println("Doing work...")
			time.Sleep(200 * time.Millisecond) // throttle so we don't burn CPU
		}
	}
}

func doneChannelExample() {
	done := make(chan struct{})

	go doWork(done)

	// Let it run for a bit, then cancel.
	time.Sleep(2 * time.Second)
	close(done) // broadcast cancellation
}

// ------------------------------------------------------------
// 4) PIPELINE (slice -> channel -> square -> print)
// ------------------------------------------------------------

// stage 1: convert slice to channel
func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		// Producer: push values then close channel
		for _, n := range nums {
			out <- n
		}
		close(out) // critical: tells downstream "no more values"
	}()

	return out
}

// stage 2: transform (square)
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		// range reads until `in` is closed
		for n := range in {
			out <- n * n
		}
		close(out) // critical: closes stage output
	}()

	return out
}

func pipelineExample() {
	nums := []int{1, 9, 3, 7, 5}

	// Stage chaining
	ch1 := sliceToChannel(nums)
	ch2 := square(ch1)

	// Final stage: consume results
	for n := range ch2 {
		fmt.Println("Squared:", n) // FIX: actually print n
	}
}

// ------------------------------------------------------------
// main (choose what to run)
// ------------------------------------------------------------

func main() {
	selectTwoChannels()
	// selectTwoChannelsSafe()

	bufferedChannelExample()

	doneChannelExample()

	pipelineExample()
}
