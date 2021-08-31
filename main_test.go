package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Notes:
//use reflect.DeepEqual for comparing whether things like maps are equal or not

func TestHandler(t *testing.T) {
	//setup
	req := httptest.NewRequest("GET", "http://example.com/home", nil)
	w := httptest.NewRecorder()
	expected := "This is go server"

	//Invoking the handler for testing
	HomeHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	t.Log(resp.StatusCode)
	t.Log(resp.Header.Get("Content-Type"))
	t.Log(string(body))

	t.Log(string(body))

	// if string(body) != "This is go server" {
	// 	t.Fatal("Failed not returning proper string")
	// }

	assert.Equal(t, string(body), expected, "Failed not returning proper string")
}
