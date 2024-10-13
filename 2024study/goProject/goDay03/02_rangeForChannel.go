package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	//遍历完成时自动跳出循环
	for num := range ch {
		fmt.Println("num = ", num)
	}
}
