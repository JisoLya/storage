package main

import "fmt"

func test01() *int {
	a := 8
	return &a
}

func main() {
	myMap := map[int]string{}

	myMap[0] = "cpp"
	myMap[1] = "java"
	myMap[3] = "go"
	fmt.Println(len(myMap))
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
