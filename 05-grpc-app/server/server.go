package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type appService struct {
	proto.UnimplementedAppServiceServer
}

// overriding the UnimplementedAppServiceServer.Add method
func (asi *appService) Add(ctx context.Context, req *proto.MathOperationRequest) (*proto.MathOperationResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Received Add request for x = %d and y = %d\n", x, y)
	result := x + y
	res := &proto.MathOperationResponse{
		Result: result,
	}
	fmt.Printf("Sending Add response with result = %d\n", result)
	return res, nil
}

func (asi *appService) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	log.Printf("Received GeneratePrimes request for start = %d and end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			log.Printf("Sending PrimeNo = %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			if err := serverStream.Send(res); err != nil {
				log.Println(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	asi := &appService{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
