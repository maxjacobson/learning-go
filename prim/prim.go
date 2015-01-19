package prim

import "math"

func IsPrime(num int) bool {
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
