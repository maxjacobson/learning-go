package main
import "fmt"

func main() {
  i := 1
  for !coolNumber(i, 20) {
    i++
  }
  fmt.Println(i)
}

func coolNumber(num int, top int) bool {
  for i:= top; i >= 1; i-- {
    if num % i != 0 {
      return false
    }
  }
  return true
}

