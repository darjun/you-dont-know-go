package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB // 2 ^ 20
	GB // 2 ^ 30
	TB // 2 ^ 40
	PB // 2 ^ 50
	EB // 2 ^ 60
	ZB // 2 ^ 70，1180591620717411303424
	YB // 2 ^ 80
)

func main() {
	fmt.Println(YB / ZB)
	fmt.Println("1180591620717411303424 B = ", 1180591620717411303424/ZB, "ZB")
}
