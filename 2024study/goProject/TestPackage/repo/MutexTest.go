package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func PrintInfo1() {
	for {
		mutex.Lock()
		fmt.Println("func1 call goroutine...")
		time.Sleep(time.Second)
		mutex.Unlock()
	}
}
func PrintInfo2() {
	for {
		mutex.Lock()
		fmt.Println("func2 call goroutine...")
		time.Sleep(time.Second)
		mutex.Unlock()
	}
}
func main() {

	go PrintInfo1()
	time.Now()
	go PrintInfo2()
	for {

	}
}
