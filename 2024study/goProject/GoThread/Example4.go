package main

import (
	"fmt"
	"time"
)

func main() {
	bools := make(chan bool, 1)

	bools <- true
	go func() {
		<-bools
	}()
	fmt.Println("go")
	time.Sleep(time.Second)
}
