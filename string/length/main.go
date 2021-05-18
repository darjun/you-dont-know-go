package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Hello World!"
	s2 := "你好，中国"

	fmt.Println(len(s1))
	fmt.Println(len(s2))

	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Println(utf8.RuneCountInString(s2))
}
