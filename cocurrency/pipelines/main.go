package main

import (
	"fmt"
	"sync"
)

func main() {
	// Set up the pipeline
	c := gen(2, 3)
	out := sq(c)

	// Consume the output
	fmt.Println(<-out)
	fmt.Println(<-out)

	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}

	// Fan-out, Fan-in
	in := gen(2, 3)

	// Distribute the sq work across two goroutines
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
