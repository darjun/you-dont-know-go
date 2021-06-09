package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(Letters[rand.Intn(len(Letters))])
	}
	return buf.String()
}

func doHTTPRequest(path string) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("ret:", len(data))
	resp.Body.Close()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			doHTTPRequest(fmt.Sprintf("http://localhost:8080/fib/%d", rand.Intn(30)))
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			doHTTPRequest(fmt.Sprintf("http://localhost:8080/repeat/%s/%d", generate(rand.Intn(200)), rand.Intn(200)))
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
}
