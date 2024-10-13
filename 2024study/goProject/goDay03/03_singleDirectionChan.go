package main

import "fmt"

func Producer(ch chan<- int) {
	//只向通道中生产数据
	for i := 0; i < 10; i++ {
		fmt.Println("Producer running ...")
		ch <- i
	}
}

func Consumer(ch <-chan int) {
	for {
		fmt.Println("Consumer running ...", <-ch)
	}
}

func main() {
	ints := make(chan int)
	go Producer(ints)
	go Consumer(ints)

	for true {

	}
}
