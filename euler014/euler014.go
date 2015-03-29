package main

import (
	"fmt"

	"github.com/maxjacobson/learning-go/collatz"
)

// https://projecteuler.net/problem=14
// Which starting number, under one million, produces the longest chain?

func main() {
	highestStartingPoint := 0
	highestValue := 0
	for i := 1; i < 1000000; i++ {
		latz := collatz.From(i)
		if latz.Length > highestValue {
			highestStartingPoint = i
			highestValue = latz.Length
		}
	}

	fmt.Println(highestStartingPoint)
}
