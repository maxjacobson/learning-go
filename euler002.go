package main

import "fmt"

func main() {
	var lastValue int
	var sum int

	for i := 1; lastValue < 4000000; i++ {
		lastValue = fibonacci(i)
		if lastValue%2 == 0 {
			sum += lastValue
		}
	}

	fmt.Println(sum)
}

func fibonacci(index int) int {
	if index < 2 {
		return 1
	} else {
		return fibonacci(index-1) + fibonacci(index-2)
	}
}
