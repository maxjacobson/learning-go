package main

import (
	"fmt"

	"github.com/maxjacobson/learning-go/prim"
)
import "math"

func main() {
	num := 600851475143
	primeFactors := []int{}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if isFactor(num, i) && prim.IsPrime(i) {
			primeFactors = append(primeFactors, i)
		}
	}
	fmt.Println(primeFactors[len(primeFactors)-1])
}

func isFactor(num int, num2 int) bool {
	return num%num2 == 0
}
