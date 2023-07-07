package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	appServiceClient := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()
	// doRequestResponse(ctx, appServiceClient)
	doServerStreaming(ctx, appServiceClient)
}

func doRequestResponse(ctx context.Context, appServiceClient proto.AppServiceClient) {
	req := &proto.MathOperationRequest{
		X: 100,
		Y: 200,
	}
	res, err := appServiceClient.Subtract(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Result :", res.GetResult())
}

func doServerStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 2,
		End:   100,
	}
	clientStream, err := appServiceClient.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		resp, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All the primes are generated!")
			break
		}
		if err != nil {
			log.Fatalln(err)
			break
		}
		fmt.Printf("Prime No : %d\n", resp.GetPrimeNo())
	}
}
