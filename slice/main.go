package main

import "fmt"

func nilEmptySlice() {
	var s1 []uint32
	s2 := make([]uint32, 0)

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println("nil slice:", len(s1), cap(s1))
	fmt.Println("cap slice:", len(s2), cap(s2))
}

func appendSlice() {
	s1 := []uint32{1, 2, 3}
	s2 := append(s1, 4)

	fmt.Println(s1)
	fmt.Println(s2)
}

func growSlice() {
	var s1 []uint32
	s1 = append(s1, 1, 2, 3)

	s2 := append(s1, 4)
	fmt.Println(&s1[0] == &s2[0])
}

func sliceString() {
	str := "hello, world"

	fmt.Println(str[:5])
}

func sliceString2() {
	str := "中国"

	fmt.Println(str[1:])
}

func sliceUpperBound() {
	array := [10]uint32{1, 2, 3, 4, 5}
	s1 := array[:5]

	s2 := s1[5:10]
	fmt.Println(s2)

	s1 = append(s1, 6)
	fmt.Println(s1)
	fmt.Println(s2)
}

func sliceUpperBound2() {
	array := [10]uint32{1, 2, 3, 4, 5}
	s1 := array[:5:5]

	s2 := array[5:10:10]
	fmt.Println(s2)

	s1 = append(s1, 6)
	fmt.Println(s1)
	fmt.Println(s2)
}

func main() {
	nilEmptySlice()

	appendSlice()

	growSlice()

	sliceString()
	sliceString2()

	sliceUpperBound()
	sliceUpperBound2()
}
