package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const (
	port = 1234
)

type HelloWorldRequest struct {
	Name string `json:"name"`
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	c := CreateClient()
	defer c.Close()

	fmt.Println(PerformRequest(c))
	fmt.Println(PerformRequest(c))

}

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func PerformRequest(client *rpc.Client) HelloWorldResponse {
	args := &HelloWorldRequest{Name: "World"}
	var reply HelloWorldResponse

	err := client.Call("HelloWorldHandler.HellWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}
