// https://projecteuler.net/problem=9
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main
import "fmt"

func main() {
  a, b, c := findBeautifulTriplet(1000)
  fmt.Println(a * b * c)
}

func findBeautifulTriplet(upperLimit int) (a, b, c int)  {
  for a := 1; a < upperLimit - 1; a++ {
    for b := 1; b < upperLimit - 1; b++ {
      c := upperLimit - b - a
      if isPythagorean(a, b, c) {
        return a, b, c
      }
    }
  }
  return 0, 0, 0 // wait, so what should be here?
}

func isPythagorean(a, b, c int) bool {
  return (a * a + b * b) == c * c
}
