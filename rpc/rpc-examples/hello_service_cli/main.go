package main

import (
	"fmt"
	"log"

	"github.com/yufeifly/advance-go/rpc-examples/hello_service/pb"
)

//type HelloServiceClient struct {
//	*rpc.Client
//}
//
//var _ service.HelloServiceInterface = (*HelloServiceClient)(nil)
//
//
//func DialHelloService(network, address string) (*HelloServiceClient, error) {
//	c, err := net.Dial(network, address)
//	if err != nil {
//		return nil, err
//	}
//	return &HelloServiceClient{Client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))}, nil
//}
//
//func (p *HelloServiceClient) Hello(request *pb.String, reply *pb.String) error {
//	return p.Client.Call("HelloService.Hello", request, reply)
//}

func main() {
	//client, err := DialHelloService("tcp", "localhost:1234")
	//if err != nil {
	//	log.Fatal("dialing:", err)
	//}
	//
	//var reply pb.String
	//var request pb.String
	//request.Value = "hello"
	//err = client.Hello(&request, &reply)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("reply: %v\n", reply)
	client, err := pb.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply pb.String
	var request pb.String
	request.Value = "hello"
	err = client.Hello(&request, &reply)
	if err != nil {
		fmt.Println("fuck")
		log.Fatal(err)
	}
	fmt.Printf("reply: %v\n", reply)

}
