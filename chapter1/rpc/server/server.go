package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const port = 1234

func main() {

}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct{}

func (h HelloWorldHandler) HellWorld(args *contract.helloWorldRequest, reply *contract.helloWorldResponse) error {
	reply.Message = "Hello" + args.Name
	return nil
}
