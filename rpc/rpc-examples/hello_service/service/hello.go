package service

import (
	"net/rpc"

	"github.com/yufeifly/advance-go/rpc-examples/hello_service/pb"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request *pb.String, reply *pb.String) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
