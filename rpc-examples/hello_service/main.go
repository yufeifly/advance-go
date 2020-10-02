package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/yufeifly/advance-go/rpc-examples/hello_service/pb"
)

type HelloService struct{}

func (p *HelloService) Hello(request *pb.String, reply *pb.String) error {
	reply.Value = "hello:" + request.Value
	return nil
}

func main() {
	//service.RegisterHelloService(new(HelloService))
	pb.RegisterHelloService(rpc.NewServer(), new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		fmt.Printf("receive con: %v\n", conn)
		//go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
