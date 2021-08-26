package util

import (
	"testing"
)

func TestSquare(t *testing.T) {
	//setup
	input := 4
	expected := 16

	//Invoke or call action
	out := Square(input)

	//Assert or verify
	if out != expected {
		t.Errorf("Failed expected %d got %d", expected, out)
	}

}

//4 -> 16
//2 -> 4
//3 -> 9
func TestSquareTabFormat(t *testing.T) {

	type test struct {
		input    int
		expected int
	}

	tests := []test{
		{input: 4, expected: 16},
		{input: 2, expected: 4},
		{input: 3, expected: 9},
		{input: 5, expected: 25},
		{input: 25, expected: 625},
	}

	for _, tc := range tests {
		out := Square(tc.input)
		if out != tc.expected {
			t.Errorf("Expected %d however got % d", tc.expected, out)
		}
	}

}

func TestSquareSubTests(t *testing.T) {

	type test struct {
		name     string
		input    int
		expected int
	}

	tests := []test{
		{name: "Eval for 4", input: 4, expected: 16},
		{name: "Eval for 2", input: 2, expected: 4},
		{name: "Eval for 3", input: 3, expected: 9},
		{name: "Eval for 5", input: 5, expected: 25},
		{name: "Eval for 25", input: 25, expected: 625},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			out := Square(tc.input)
			if out != tc.expected {
				t.Fatalf("Expected %d however got % d", tc.expected, out)
			}
		})
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(20)
	}
}
