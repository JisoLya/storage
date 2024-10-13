package main

import (
	"fmt"
	"time"
)

func main() {
	//无缓冲的channel
	ch := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			ch <- i
			time.Sleep(2 * time.Second)
		}
		//没有数据要写时，关闭channel
		close(ch)
	}()

	for { //分别是数据和管道的状态
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else {
			break
		}
	}
}
