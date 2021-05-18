package main

import "testing"

func BenchmarkConcatWithOnePlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatWithOnePlus()
	}
}

func BenchmarkConcatWithMultiPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatWithMultiPlus()
	}
}

func BenchmarkConcatWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatWithJoin()
	}
}
