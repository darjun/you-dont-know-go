package main

import "fmt"

// const PI float64 = 3.1415926
// const MaxAge int = 150
// const Greeting string = "hello world"

const (
	PI       float64 = 3.1415926
	MaxAge   int     = 150
	Greeting string  = "hello world"
)

const b byte = 128
const r rune = 'c'

// type User struct {
// 	Name string
// 	Age  int
// }

// const u User = User{}

// var i int = 1

// const p *int = &i

func main() {
	fmt.Println(PI)
	fmt.Println(MaxAge)
	fmt.Println(Greeting)

	fmt.Println(b)
	fmt.Println(r)
}
