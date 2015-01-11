// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"math"
)

type Prime struct {
	num     int
	isPrime bool
}

func main() {
	sum := 0
	upper_limit := 2000000
	c := primes()
	for i := 2; i < upper_limit; i++ {
		prime := <-c
		if prime.isPrime {
			sum += prime.num
		}
	}
	fmt.Println(sum)
}

func primes() <-chan Prime {
	c := make(chan Prime)
	go func() {
		for i := 2; ; i++ {
			c <- Prime{
				num:     i,
				isPrime: isPrime(i),
			}
		}
	}()
	return c
}

func isPrime(num int) bool {
	magic := sqrt(num)
	for i := 2; i <= magic; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func sqrt(num int) int {
	return int(math.Sqrt(float64(num)))
}
