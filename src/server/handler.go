package server

import (
	"fmt"
	"net/http"
	"sync"
)

type Handler func(http.ResponseWriter, *http.Request)

var handlers map[string]Handler
var mu sync.RWMutex

func init() {
	handlers = make(map[string]Handler)
}

func registerHandler(name string, handler Handler) {
	mu.Lock()
	defer mu.Unlock()
	handlers[name] = handler
}

func server() {
	http.HandleFunc("/", handleRoot)
	for name := range handlers {
		http.HandleFunc("/"+name, handlers[name])
	}
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my server!")
}
