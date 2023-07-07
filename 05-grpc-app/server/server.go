package main

import (
	"fmt"
	"grpc-app/proto"
	services "grpc-app/server/services/appService"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	asi := &services.AppService{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	go func() {
		fmt.Scanln()
		fmt.Println("Stopping server....")
		grpcServer.Stop()
	}()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
