package main

import (
	"fmt"
	"strings"
	"testing"
)

const (
	// ExampleInput : This is a sample input sentence
	ExampleInput = "Live your life like there is no tomorrow." 

	// ExampleOutput : This is a sample reverse output sentence
	ExampleOutput = "tomorrow. no is there like life your Live" 

	// ErrorResult : This is a error string
	ErrorResult = "error, unexpected result: %v"

	// SingleSpace : Single space string
	SingleSpace = " "

)

// BenchmarkStringReverseBad : This function is a 
// bad way of benchmarking a string reverse function
func BenchmarkStringReverseBad(b *testing.B) {
	b.ReportAllocs()

	

	for i := 0; i<b.N; i++ {
		words := strings.Split(ExampleInput, SingleSpace)
		wordsReverse := make([]string, 0)

		for {
			word := words[len(words) - 1:][0]
			wordsReverse = append(wordsReverse, word)
			words = words[:len(words) - 1]
			if len(words) == 0 {
				break
			}
		}

		output := strings.Join(wordsReverse, SingleSpace)
		if output != ExampleOutput{
			b.Error(fmt.Sprintf(ErrorResult, output))
		}
	}
}

// BenchmarkStringReverseBetter : This is a
// good way of benchmarking a string reverse function
func BenchmarkStringReverseBetter(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i<b.N; i ++ {

		words := strings.Split(ExampleInput, SingleSpace)
		
		for i := 0; i < len(words)/2; i++ {
			words[len(words) -1 -i], words[i] = words[i], words[len(words)-1-i]
		}

		output := strings.Join(words, SingleSpace)
		
		if output != ExampleOutput {
			b.Error(fmt.Sprintf(ErrorResult, output))
		}
	}
}