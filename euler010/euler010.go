// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import (
	"fmt"

	"github.com/maxjacobson/learning-go/prim"
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
				isPrime: prim.IsPrime(i),
			}
		}
	}()
	return c
}
