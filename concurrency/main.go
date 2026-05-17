package main

import (
	"fmt"
	"sync"
	"time"
)

//
// BASIC FUNCTION
//

func greet() {
	fmt.Println("hello from greet")
}

//
// SIMPLE COUNT FUNCTION
//

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

//
// CHANNEL COUNT FUNCTION
//

func channelCount(name string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- name
		time.Sleep(500 * time.Millisecond)
	}

	close(c)
}

//
// WORKER POOL
//

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {

		fmt.Println("worker", id, "started job", j)

		time.Sleep(time.Second)

		fmt.Println("worker", id, "finished job", j)

		results <- fib(j)
	}
}

//
// FIBONACCI
//

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {

	//
	// NORMAL FUNCTION
	//

	greet()

	//
	// SINGLE GOROUTINE
	//

	go count("sheep")

	time.Sleep(3 * time.Second)

	//
	// MULTIPLE GOROUTINES
	//

	go count("fish")
	go count("cow")

	time.Sleep(3 * time.Second)

	//
	// WAITGROUP
	//

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		count("waitgroup")
		wg.Done()
	}()

	wg.Wait()

	//
	// BASIC CHANNEL
	//

	c := make(chan string)

	go func() {
		c <- "hello from channel"
	}()

	msg := <-c

	fmt.Println(msg)

	//
	// BUFFERED CHANNEL
	//

	buffered := make(chan string, 2)

	buffered <- "first"
	buffered <- "second"

	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	//
	// CHANNEL LOOPING
	//

	channel := make(chan string)

	go channelCount("sheep", channel)

	for msg := range channel {
		fmt.Println(msg)
	}

	//
	// WORKER POOL
	//

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)
	go worker(4, jobs, results)

	for i := 0; i < 10; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 10; j++ {
		fmt.Println("fib result:", <-results)
	}
}
