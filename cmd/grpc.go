package cmd

import (
	"ewallet-ums/helpers"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		logrus.Fatal("failed to listen grpc port: ", err)
	}

	s := grpc.NewServer()

	//list method
	// pb.ExampleMethod(s, &grpc....)

	logrus.Info("start listening grpc on port: " + helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(lis); err != nil {
		logrus.Fatal("failed to serve grpc port: ", err)
	}
}
