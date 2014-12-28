package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func (person *Person) Greet() {
  fmt.Println("Hello, my name is", person.Name, "and I am", person.Age)
}

func main() {
  max := Person{
    Name: "Max",
    Age: 26,
  }

  greet(max)

  max.Greet()
}

func greet(person Person) {
  fmt.Println("Hello, my name is", person.Name, "and I am", person.Age)
}
