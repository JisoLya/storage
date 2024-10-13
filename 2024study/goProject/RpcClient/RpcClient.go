package main

import (
	"fmt"
	"sync"
)

func main() {
	var map0 sync.Map
	map0.Store("123", 1)
	map0.Store("234", 1)
	fmt.Println(map0.Load("123"))
	map0.Store("123", 2)
	fmt.Println(map0.Load("123"))
}
