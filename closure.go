package main

import "fmt"

func Adder() func(int64) int64 {
	sum := int64(0)

	return func(i int64) int64 {
		sum += i
		return sum
	}
}

func MakeAdder(a int64) func(int64) int64 {
	return func(b int64) int64 {
		return a + b
	}
}

func SimpleClosureDemo() {
	adder := Adder()
	fmt.Println(adder(5))
	fmt.Println(adder(5))

	adder1 := Adder()
	fmt.Println(adder1(5))
}

func ClosureDemoWithAdder() {
	fiveAdder := MakeAdder(5)
	fmt.Println(fiveAdder(1))

	twoAdder := MakeAdder(2)
	fmt.Println(twoAdder(2))
}
