package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type PersonService struct {
	name    string
	age     uint8
	address string
}
type TaskArgs struct {
}

func (p *PersonService) CreatePerson(args *TaskArgs, reply *PersonService) error {
	reply.name = "Mike"
	reply.age = 23
	reply.address = "HomeLand"
	fmt.Println(args)
	return nil
}

func main() {
	gob.Register(TaskArgs{})
	err := rpc.Register(new(PersonService))
	if err != nil {
		log.Fatal("Registering PersonService failed!")
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		//
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			//
		}
		go rpc.ServeConn(conn)
	}
}
