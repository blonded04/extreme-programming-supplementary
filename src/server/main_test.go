package server

import (
	"net/http"
	"testing"
	"time"
)

func startServer() {
	registerHandler("/api/get_user_data", getHandler)
	go server()
	time.Sleep(time.Second)
}

func TestServerStarts(t *testing.T) {
	startServer()

	resp, err := http.Get("http://localhost:8080/api/get_user_data")
	if err != nil {
		t.Fatalf("Не удалось сделать запрос к серверу: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Ожидался статус 200 OK, но получили: %d", resp.StatusCode)
	}
}
