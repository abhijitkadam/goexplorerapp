package main

import (
	"context"
	"fmt"
	"net/http"
)

func BasicAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || !AuthCheck(user, pass) {
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(rw, "Not Authorized")
			return
		}

		context := context.WithValue(r.Context(), "authuser", user)
		h(rw, r.WithContext(context))
	}
}

//AuthCheck should connect with DB to authenticate the user
func AuthCheck(user, pass string) bool {

	if user == "abc" && pass == "xyz" {
		return true
	}

	return false
}
