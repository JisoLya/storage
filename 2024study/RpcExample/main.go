package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Person struct {
	Name    string
	Age     uint8
	Address string
	Id      int
}
type TaskArgs struct {
}
type Task struct {
}

func getPerson() Person {
	person := Person{}
	args := TaskArgs{}
	ok := call2("PersonService.CreatePerson", &args, &person)
	if !ok {
		log.Println("call2 failed")
	}
	return person
}

func call2(rpcname string, args interface{}, reply interface{}) bool {
	//person := Person{}
	//args1 := TaskArgs{}
	client, err := rpc.Dial("tcp", ":8081")
	if err != nil {
		log.Println("Dial tcp :8081 failed!")
		return false
	}
	err = client.Call(rpcname, args, reply)
	if err != nil {
		log.Println("Call method failed!")
		log.Println(err)
		return false
	}
	//reply = &person

	time.Sleep(time.Second)
	return true
}

func call(args *Person) Person {
	var person Person
	client, err := rpc.Dial("tcp", ":8081")
	if err != nil {
		log.Println("Dial tcp :8081 failed!")
	}
	err = client.Call("PersonService.CreatePerson", args, &person)
	if err != nil {
		log.Println("Call method failed!")
		log.Println(err)
	}
	//person = *reply
	return person
}

func main() {

	person := Person{
		Name:    "Rose",
		Age:     20,
		Address: "China",
		Id:      123,
	}
	m := make(map[int]*Person, 3)
	m[0] = &person
	people := make(chan *Person, 3)
	people <- &person
	person.Name = "Jisoo"
	fmt.Println(<-people)
	value, ok := m[1]
	fmt.Println(value, ok)
	//fmt.Println(person)
	//fmt.Println(receive)
}
