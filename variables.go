package main

import "fmt"

func main() {
  var message string
  message = "Hello World"
  fmt.Println(message)

  message2 := "Message with type inferred"

  fmt.Println(message2)

  var number int
  if number == 0 {
    fmt.Println("ints default to 0")
  }

  var str string
  if str == "" {
    fmt.Println("strings default to ''")
  }

  var boo bool
  if boo == false {
    fmt.Println("bools default to false")
  }

  var boo2 = false
  fmt.Println(boo2)
}

