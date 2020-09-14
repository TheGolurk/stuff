package main

import "fmt"

func main() {
	fib(11000, 1, 1)
}

// fibonacci 1 1 2 3 5 .....
func fib(number, prev, next uint64) {
	if number == 0 {
		return
	}
	flag := prev
	prev = next
	next += flag
	fmt.Printf("|%d %d = %d | \n", flag, prev, next)
	fib(number-1, prev, next)
}
