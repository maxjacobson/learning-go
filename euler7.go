package main
import "fmt"
import "math"

func main() {
  goal := 10001
  primes := []int{}
  for i := 2; len(primes) < goal; i++ {
    if isPrime(i) {
      primes = append(primes, i)
    }
  }
  fmt.Println(primes[goal - 1])
}

func isPrime(num int) bool {
  lowestPossible := int( math.Sqrt(float64(num)) )
  if lowestPossible == 1 {
    return true
  }
  for i := lowestPossible; i < num; i++ {
    if num % i == 0 {
      return false
    }
  }
  return true
}
