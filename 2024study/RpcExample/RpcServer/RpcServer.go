package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type PersonService struct {
	Name    string
	Age     uint8
	Address string
	testMap map[int]string
}
type Person struct {
	Name    string
	Age     uint8
	Address string
	Id      int
}
type TaskArgs struct {
}

func (p *PersonService) CreatePerson(args Person, reply *Person) error {
	fmt.Println("args = ", args)
	args.Name = "Jisoo"
	*reply = args
	fmt.Println("reply =", reply)
	return nil
}

func (p *PersonService) Test(args []interface{}, reply *Person) error {
	fmt.Println("Call success!")

	for _, value := range args {
		fmt.Printf("%T", value)
		fmt.Println(value)
	}
	return nil
}

func (p *PersonService) ModifyMap(args Person, reply *Person) error {
	p.testMap[args.Id] = "aaa"
	fmt.Println(p.testMap)
	str := p.testMap[args.Id]
	str = "364"
	fmt.Println(str, p.testMap)
	return nil
}

func main() {
	personService := PersonService{
		Name:    "123",
		Age:     12,
		Address: "111",
		testMap: make(map[int]string, 3),
	}

	log.Println("RpcServer.Main is running!")
	err := rpc.Register(&personService)
	if err != nil {
		log.Fatal("Registering PersonService failed!")
	}
	rpc.HandleHTTP()
	log.Println("Successfully register service")
	listener, err := net.Listen("tcp", ":8081")
	log.Println("Server is listening on port : 8081")
	if err != nil {
	}
	//conn, err := listener.Accept()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
