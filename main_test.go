package main

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	//setup
	req := httptest.NewRequest("GET", "http://example.com/home", nil)
	w := httptest.NewRecorder()

	//Invoking the handler for testing
	HomeHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if string(body) != "This is go server" {
		t.Fatal("Failed not returning proper string")
	}
}
