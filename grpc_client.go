package boom

import (
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

// ClientConn...
type ClientGrpcConn struct {
	Name  string
	Port  string
	Creds *credentials.TransportCredentials
	Kacp  *keepalive.ClientParameters
}

// NewClientGrpc...
func NewClientGrpc(opt ClientGrpcConn) (*grpc.ClientConn, error) {
	// If not implemented keepalive
	if opt.Kacp == nil {
		opt.Kacp = &keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             time.Second,
			PermitWithoutStream: true,
		}
	}

	// If not exist crentials
	if opt.Creds == nil {
		creds := insecure.NewCredentials()
		opt.Creds = &creds
	}

	conn, err := grpc.Dial(opt.Port, grpc.WithTransportCredentials(*opt.Creds), grpc.WithKeepaliveParams(*opt.Kacp))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
		return nil, err
	}

	logrus.Printf("start grpc client %s at host %s", opt.Name, opt.Port)
	return conn, nil
}
