package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
)

func NewProfileHttpServer(addr string) {
	go func() {
		log.Fatalln(http.ListenAndServe(addr, nil))
	}()
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}

	return result
}

func response(w http.ResponseWriter, data interface{}) {
	m := map[string]interface{}{
		"data": data,
	}
	bytes, _ := json.Marshal(m)
	w.Write(bytes)
}

func responseError(w http.ResponseWriter, err error) {
	m := map[string]interface{}{
		"error": err.Error(),
	}
	bytes, _ := json.Marshal(m)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(bytes)
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Path[len("/fib/"):])
	if err != nil {
		responseError(w, err)
		return
	}

	var result int
	for i := 0; i < 1000; i++ {
		result = fib(n)
	}
	response(w, result)
}

func repeatHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.SplitN(r.URL.Path[len("/repeat/"):], "/", 2)
	if len(parts) != 2 {
		responseError(w, errors.New("invalid params"))
		return
	}

	s := parts[0]
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		responseError(w, err)
		return
	}

	var result string
	for i := 0; i < 1000; i++ {
		result = repeat(s, n)
	}
	response(w, result)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fib/", fibHandler)
	mux.HandleFunc("/repeat/", repeatHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	NewProfileHttpServer(":9999")

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
