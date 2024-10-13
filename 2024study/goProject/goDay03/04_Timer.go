package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，时间为2s,2s后向time通道中写内容(当前时间)
	//只会向通道中写一次数据
	timer := time.NewTimer(time.Second)
	fmt.Println("current time : ", time.Now())
	//t := <-timer.C //channel没有数据前会阻塞
	//fmt.Println("t = ", t)

	for {
		fmt.Println(<-timer.C) // 导致死锁
	}
}
