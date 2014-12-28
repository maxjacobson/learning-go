package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func main() {
  max := Person{
    Name: "Max",
    Age: 26,
  }

  greet(max)
}

func greet(person Person) {
  fmt.Println("Hello, my name is", person.Name, "and I am", person.Age)
}

