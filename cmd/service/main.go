package main

import (
	"flag"
	"fmt"
	"github.com/fatelei/grpc-svn/config"
	proto "github.com/fatelei/grpc-svn/internal"
	"github.com/fatelei/grpc-svn/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int64("port", 8000, "-port")
	configPath := flag.String("config", "/etc/grpc_svn.toml", "-config")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	config.ParseConfig(configPath)
	grpcServer := grpc.NewServer()
	proto.RegisterGrpcSvnServiceServer(grpcServer, service.NewGrpcSvnServiceServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
