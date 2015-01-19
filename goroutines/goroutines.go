package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("started")
	go process()
	time.Sleep(time.Millisecond)
	fmt.Println("done")
}

func process() {
	fmt.Println("processing...")
}
