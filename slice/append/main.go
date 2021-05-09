package main

import "fmt"

func main() {
	s := make([]uint32, 0, 4)

	s = append(s, 1, 2, 3)
	fmt.Println(len(s), cap(s))

	s = append(s, 4, 5, 6)
	fmt.Println(len(s), cap(s))
}
