package main

import "fmt"

func main() {
	// top := 10
	top := 100
	diff := squareOfSum(top) - sumOfSquares(top)
	fmt.Println(diff)
}

func squareOfSum(top int) int {
	var sum int
	for i := 1; i <= top; i++ {
		sum += i
	}
	return sum * sum
}

func sumOfSquares(top int) int {
	var sum int
	for i := 1; i <= top; i++ {
		sum += i * i
	}
	return sum
}
