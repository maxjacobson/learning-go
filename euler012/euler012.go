package main

import (
	"fmt"

	"github.com/maxjacobson/learning-go/divisors"
	"github.com/maxjacobson/learning-go/trinums"
)

func main() {
	for i := 1; ; i++ {
		num := trinums.At(i)
		div := divisors.New(num)
		if div.MoreThan(500) {
			fmt.Println(num)
			break
		}
	}
}
