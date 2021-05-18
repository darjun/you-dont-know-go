package main

import "fmt"

func main() {
	s := "中国"
	fmt.Println(s[:5])

	b := []byte{129, 130, 131}
	fmt.Println(string(b))
}
