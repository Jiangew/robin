package main

import (
	"fmt"
	"testing"
)

// go test
func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := 4
	result := Calculate(2)
	if result != expected {
		t.Error("Failed, excepted '2 + 2' to equal 4")
	}
}

// Table Driven Testing: go test -v
func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{999999, 1000001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This shouldn't run with -run=calc")
}

// Simple Benchmark Test: go test -bench=.
// Benchmark with -run Flag
// If we wanted to just run our Calculate tests, we could specify a regex that matches on Calculate: go test -run=Calculate -bench=.
// If we wanted to only run our BenchmarkCalculate function then we could change our regex pattern to be: go test -run=Bench -bench=.
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2)
	}
}

func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

// Benchmark Test: go test -run=Bench -bench=.
func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
