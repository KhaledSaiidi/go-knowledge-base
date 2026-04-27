package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wgForChannel *sync.WaitGroup

func doWorkWithWg(t time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing work...")
	time.Sleep(t)
	fmt.Println("work done!")
}

func doWorkWithChans(
	t time.Duration,
	resultChannel chan string,
) {
	defer wgForChannel.Done()
	fmt.Println("doing work...")
	time.Sleep(t)
	fmt.Println("work done!")
	resultChannel <- fmt.Sprintf("word %d", rand.Intn(100))
}

func main() {
	// do Work With WaitGroups
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("--------------------------------Do Work With WaitGroups--------------------------------")
	fmt.Println("---------------------------------------------------------------------------------------")
	wg := sync.WaitGroup{}
	start := time.Now()
	timeToWait := []int{1, 2}
	wg.Add(len(timeToWait))
	for i := range timeToWait {
		go doWorkWithWg(time.Second*time.Duration(timeToWait[i]), &wg)
	}
	wg.Wait()
	fmt.Printf("Time spent in doWorkWithWg is: %v\n", time.Since(start))

	// do Work With Chans
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("---------------------------------Do Work With Channels---------------------------------")
	fmt.Println("---------------------------------------------------------------------------------------")
	startWork2 := time.Now()
	resultChannel := make(chan string)
	wgForChannel = &sync.WaitGroup{}
	timeToWaitForChannel := []time.Duration{1, 3, 2}
	wgForChannel.Add(len(timeToWaitForChannel)) // This WaitGroup should track only the goroutines that produce/send data
	for i := range timeToWaitForChannel {
		go doWorkWithChans(time.Second*timeToWaitForChannel[i], resultChannel)
	}
	// res1 := <-resultChannel
	// fmt.Println(res1)
	// res2 := <-resultChannel
	// fmt.Println(res2)
	go func() {
		wgForChannel.Wait()
		close(resultChannel)
	}()

	for res := range resultChannel {
		fmt.Println(res)
	}

	fmt.Printf("Time spent in doWorkWithChans is: %v\n", time.Since(startWork2))
}
