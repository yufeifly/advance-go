package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/yufeifly/advance-go/rpc-examples/hello_service/service"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ .HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	service.HelloServiceInterface()
	//c, err := rpc.Dial(network, address)
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(service.HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reply: %v\n", reply)
}
