package boom

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerGrpcConn struct {
	Name               string
	Port               string
	ImplRegisterServer func(*grpc.Server, interface{})
	ImplService        interface{}
}

// NewServerGrpc ...
func NewServerGrpc(opt ServerGrpcConn) {
	listener := getNetListener(opt.Port)

	server := grpc.NewServer()

	reflection.Register(server)

	// Implementação do register do cliente
	opt.ImplRegisterServer(server, opt.ImplService)

	logrus.Printf("GRPC server listening at %s name server %s", opt.Port, opt.Name)
	if err := server.Serve(listener); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
		panic(err)
	}
}

func getNetListener(port string) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
