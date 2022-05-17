package boom

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
)

func ServerGRPC(server *API) {
	if server.GrpcServer == nil {
		return
	}
	reflection.Register(server.GrpcServer)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s", server.Settings.Grpc.Host))
	if err != nil {
		logrus.Fatalf("error binding address %s: %v", server.Settings.Grpc.Host, err)
	}
	go func() {
		logrus.Infof("GRPC server listening at %s", server.Settings.Grpc.Host)
		if err := server.GrpcServer.Serve(listener); err != nil {
			logrus.Fatalf("failed to serve: %v", err)
		}
	}()
}

func StopGRPC(ctx context.Context, server *API) {
	if server.GrpcServer == nil {
		return
	}
	stopChan := make(chan interface{})
	go func() {
		server.GrpcServer.GracefulStop()
		stopChan <- nil
	}()
	select {
	case <-ctx.Done():
		logrus.Infof("Error gracefully stopping GRPC server:%v", ctx.Err())
		server.GrpcServer.Stop()
		logrus.Info("GRPC server forcefully stopped")
	case <-stopChan:
		logrus.Info("GRPC server stopped gracefully")
	}
}
