package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func printSlice() {
	s := make([]uint32, 1, 10)
	fmt.Printf("%#v\n", *(*slice)(unsafe.Pointer(&s)))
}

func main() {
	printSlice()
}
