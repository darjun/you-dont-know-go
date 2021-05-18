package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello" + " " + "World"
	fmt.Println(s1)

	ss := []string{"Hello", "World"}
	fmt.Println(strings.Join(ss, " "))
}
