// select关键字
// 每一个case必须是一个IO操作，
// select {
// case <-chan1:
//
//	chan1读到数据，则进行该case
//
// case chan2 <- 1:
//
//	chan2写入数据，则进行case处理
//
// 如果case的条件都满足，那么会挑选任意一个语句执行
// 如果没有一条语句可以执行，那么会执行default
// 如果没有default语句，那么所有的select语句都会被阻塞
// }
package main

import "fmt"

// Fib 案例：select实现斐波那契数列
func Fib(ch chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}
func main() {
	ch1 := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch1
			fmt.Println(num)
		}
		quit <- true
	}()

	Fib(ch1, quit)
}
