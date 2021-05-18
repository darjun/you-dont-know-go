package main

import "fmt"

func main() {
	s := "Hello"

	b := []byte(s)
	fmt.Println(len(b), cap(b))
}
