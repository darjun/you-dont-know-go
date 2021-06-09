package main

import (
	"bytes"
	"math/rand"
	"os"
	"runtime/pprof"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(Letters[rand.Intn(len(Letters))])
	}
	return buf.String()
}

func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}

	return result
}

func main() {
	f, _ := os.OpenFile("mem.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	for i := 0; i < 100; i++ {
		repeat(generate(100), 100)
	}

	pprof.Lookup("heap").WriteTo(f, 0)
}
