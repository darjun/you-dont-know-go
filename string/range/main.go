package main

import "fmt"

func main() {
	s1 := "Go 语言"
	for index, c := range s1 {
		fmt.Println(index, c)
	}

	s2 := "Go 语言"
	for index, c := range s2 {
		fmt.Printf("%d:%c\n", index, c)
	}

	fmt.Println(Utf8Count("中国"))
}

func Utf8Count(s string) int {
	var count int
	for range s {
		count++
	}
	return count
}
