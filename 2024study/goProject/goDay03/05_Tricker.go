package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	var i = 0
	for {
		t := <-ticker.C
		i++
		fmt.Println("i = ", i, "t = ", t, time.Now())
	}
}
