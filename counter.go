package main

type Counter interface {
	Increment()
	Decrement()
	GetValue() int64
}
