package main

import (
	"fmt"
	"time"
)

func GoKeyWordDemo() {

	go SomeSimpleFunction()

	go func() {
		fmt.Println("You are inside of a func(){}")
	}()

	sfnValue := func() {
		fmt.Println("You are inside simple fn defined as function value")
	}

	go sfnValue()

	time.Sleep(time.Second * 1)
}

func SomeSimpleFunction() {
	fmt.Println("You are inside a simple function")
}
