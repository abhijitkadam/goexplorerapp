package main

type SimpleCounter struct {
	Value int64
}

func NewSimpleCounter() SimpleCounter {
	return SimpleCounter{}
}

func (sc *SimpleCounter) Increment() {
	sc.Value += 1
}

/*
func Incr(sc *SimpleCounter) {
	sc.Value += 1
}
*/

func (sc *SimpleCounter) Decrement() {
	sc.Value -= 1
}

func (sc *SimpleCounter) GetValue() int64 {
	return sc.Value
}
