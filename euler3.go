package main
import "fmt"
import "math"

func main() {
  num := 600851475143
  primeFactors := []int{}
  for i := 2; i <= int( math.Sqrt(float64(num)) ); i++ {
    if isFactor(num, i) && isPrime(i) {
      primeFactors = append(primeFactors, i)
    }
  }
  fmt.Println(primeFactors[ len(primeFactors) - 1 ])
}

func isPrime(num int) bool {
  lowestPossible := int( math.Sqrt(float64(num)) )
  for i := lowestPossible; i < num; i++ {
    if num % i == 0 {
      return false
    }
  }
  return true
}

func isFactor(num int, num2 int) bool {
  return num % num2 == 0
}

