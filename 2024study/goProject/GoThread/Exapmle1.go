package main

import (
	"sync"
	"time"
)

var locker sync.Mutex
var number int

func main() {
	newCond := sync.NewCond(&locker)
	go func() {
		locker.Lock()
		for number < 10 {
			newCond.Wait()
		}
		locker.Unlock()
		println(11)
	}()

	go func() {
		locker.Lock()
		newCond.Wait()
		locker.Unlock()
		println(12)
	}()
	a(newCond)
	time.Sleep(3 * time.Second)
	number = 20
	a(newCond)
	time.Sleep(3 * time.Second)

}

func a(cond *sync.Cond) {
	cond.Broadcast()
}
