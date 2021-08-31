package main

import (
	"context"
	"fmt"
	"net/http"
)

func MuxloggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Requesting ", r.URL)

		user := r.Context().Value("authuser")
		if user != nil {
			userName, ok := user.(string)
			if ok {
				fmt.Println("Authenticated & user is : ", userName)
			}
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func MuxbasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || !AuthCheck(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Not Authorized")
			return
		}

		context := context.WithValue(r.Context(), "authuser", user)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(context))
	})
}
