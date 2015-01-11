package main

import "fmt"

func main() {
	nums := []string{"hello", "world"}

	for index, value := range nums {
		fmt.Println(index)
		fmt.Println(value)
	}
}
