package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	defer profile.Start().Stop()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Printf("fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
}
