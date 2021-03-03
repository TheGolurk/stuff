package main

import (
	"log"
	"net/http"
	"sync"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI

	switch path {
	case "/hello":
		go sayHello(w, r)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found 😞"))
	}

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handler{})

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(nb)
	guard := make(chan struct{}, maxGoroutines)
	go func() {
		w.Write([]byte("Hello"))
	}()
	wg.Wait()
}
