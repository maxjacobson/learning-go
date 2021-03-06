package main

import "time"

var c = make(chan int, 5)

func main() {
	go worker(1)
	for i := 0; i < 10; i++ {
		c <- i
		println(i)
	}
}

func worker(id int) {
	for {
		_ = <-c
		time.Sleep(time.Second)
	}
}
