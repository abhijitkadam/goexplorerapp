package main

import "fmt"

//https://blog.golang.org/json
//https://pkg.go.dev/encoding/json#Marshal

//omitempty -> use when you don't want to include anyfield in json
//The problem with omitempty is it takes the default value I guess.
//The solution is to use pointer. So empty value of pointer will be nil.

func main() {
	//GoSlices()
	//GoSlicesRange()

	//err in go
	//Always check for error first.
	hotels, err := FetchHotelsFromUrlUsingJsonDecoder("http://localhost:3000/inventory")

	if err != nil {
		panic(err)
	}

	fmt.Println(hotels)
}
