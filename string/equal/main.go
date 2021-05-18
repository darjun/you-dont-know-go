package main

import "fmt"

func main() {
	b1 := []byte{0xEF, 0xBB, 0xBF, 72, 101, 108, 108, 111}
	b2 := []byte{72, 101, 108, 108, 111}

	s1 := string(b1)
	s2 := string(b2)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}
