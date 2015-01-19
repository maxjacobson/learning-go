package main

import (
	"fmt"

	"github.com/maxjacobson/learning-go/prim"
)

func main() {
	goal := 10001
	primes := []int{}
	for i := 2; len(primes) < goal; i++ {
		if prim.IsPrime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes[goal-1])
}
