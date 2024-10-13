package main

import "fmt"

type KeyValue struct {
	Key   string
	Value string
}

var V any

func main() {
	x := []KeyValue{}
	fmt.Printf("%T", x)
	x = append(x, KeyValue{"key", "value"})
	for a, b := range x {
		fmt.Println("a = ", a, "b = ", b)
	}
	fmt.Printf("%T", V)
}
