package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	//延时两秒后打印一句话
	//<-timer.C
	//fmt.Println("时间到")
	//重新定时
	timer.Reset(time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，定时器时间到")
	}()
	//加入这段代码后，定时器停止，子协程一直阻塞
	timer.Stop()

	for {

	}
}
