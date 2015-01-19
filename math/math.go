package main

import "math"
import "fmt"

func main() {
	var five int = 5
	num := math.Sqrt(float64(five))
	// num2 := math.Floor(num)
	var num2 int = int(num)
	fmt.Println(num2)
}
