package main

import (
	"fmt"
	"time"
)

func TestSelect(channel chan int) {
	x := 1
	for {
		select {
		case channel <- x:
			fmt.Println("向channel写入了数据 x = ", x)
			fmt.Println("此时的channel ：", channel)
			x++
			time.Sleep(5 * time.Second)
		case i := <-channel:
			fmt.Println("从channel读出了数据 i = ", i)
			fmt.Println("此时的channel ：", channel)
		}
	}
}

func main() {
	channel := make(chan int, 4)
	TestSelect(channel)
}
