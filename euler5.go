package main
import "fmt"

func main() {
  i := 1
  for !coolNumber(i) {
    i++
  }
  fmt.Println(i)
}

func coolNumber(num int) bool {
  for i := 1; i <= 20; i++ {
    if num % i != 0 {
      return false
    }
  }
  return true
}

