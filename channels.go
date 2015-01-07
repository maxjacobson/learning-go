package main

import (
  "fmt"
  "time"
)

func main() {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    worker := &Worker{id: i}
    go worker.process(c)
  }
  for i := 2; i < 10; i++ {
    c <- i
    time.Sleep(time.Millisecond * 50)
  }
}

type Worker struct {
  id int
}

func (w *Worker) process(c chan int) bool {
  fmt.Println("Hey! I'm processing! I'm", w.id)
  for {
    data := <-c
    fmt.Printf("worker %d got %d\n", w.id, data)
  }
}

// func isPrime(num int) bool {
//   lowestPossible := int( math.Sqrt(float64(num)) )
//   if lowestPossible == 1 {
//     return true
//   }
//   for i := lowestPossible; i < num; i++ {
//     if num % i == 0 {
//       return false
//     }
//   }
//   return true
// }
