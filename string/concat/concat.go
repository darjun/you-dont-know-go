package main

import "strings"

func ConcatWithMultiPlus() {
	var s string
	for i := 0; i < 10; i++ {
		s += "hello"
	}
}

func ConcatWithOnePlus() {
	s1 := "hello"
	s2 := "hello"
	s3 := "hello"
	s4 := "hello"
	s5 := "hello"
	s6 := "hello"
	s7 := "hello"
	s8 := "hello"
	s9 := "hello"
	s10 := "hello"
	s := s1 + s2 + s3 + s4 + s5 + s6 + s7 + s8 + s9 + s10
	_ = s
}

func ConcatWithJoin() {
	s := []string{"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello"}
	_ = strings.Join(s, "")
}
