package main

import "fmt"

const (
	One int = iota + 1
	Two
	Three
	Four
	Five
)

const (
	Mask1 int = 1<<(iota+1) - 1
	Mask2
	Mask3
	Mask4
)

const (
	Odd1 = 2*iota + 1
	Odd2
	Odd3
)

const (
	Even1 = 2 * (iota + 1)
	Even2
	Even3
)

const (
	A int = 1
	B int = 2
	C int = iota + 1
	D
	E
)

func main() {
	fmt.Println(One, Two, Three, Four, Five)
	fmt.Println(Mask1, Mask2, Mask3, Mask4)
	fmt.Println(Odd1, Odd2, Odd3)
	fmt.Println(Even1, Even2, Even3)
	fmt.Println(A, B, C, D, E)
}
