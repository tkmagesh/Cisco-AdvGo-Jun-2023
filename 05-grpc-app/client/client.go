package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	// doServerStreaming(ctx, appServiceClient)
	// doClientStreaming(ctx, appServiceClient)
	/* doneCh := doBiDiStreaming(ctx, appServiceClient)
	<-doneCh */
	doRequestResponseWithTimeout(ctx, appServiceClient)
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

func doRequestResponseWithTimeout(ctx context.Context, appServiceClient proto.AppServiceClient) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	addReq := &proto.MathOperationRequest{
		X: 100,
		Y: 200,
	}

	res, err := appServiceClient.Add(timeoutCtx, addReq)
	if err != nil {
		if code := status.Code(err); code == codes.DeadlineExceeded {
			fmt.Println("timeout occurred")
			return
		}
		log.Fatalln(err)
	}
	fmt.Printf("Result = %d\n", res.GetResult())
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
		time.Sleep(500 * time.Millisecond)
	}
}

func doClientStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	data := []int32{3, 1, 4, 2, 5, 6, 8, 7, 9}
	clientStream, err := appServiceClient.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range data {
		req := &proto.AverageRequest{
			No: no,
		}
		log.Printf("Sending No : %d\n", no)
		clientStream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Average : %d\n", res.GetAverage())
}

func doBiDiStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) <-chan struct{} {
	clientStream, err := appServiceClient.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	go sendRequests(ctx, clientStream)
	done := make(chan struct{})
	go func() {
		fmt.Println("Press ENTER to cancel")
		fmt.Scanln()
		clientStream.CloseSend()
		close(done)
	}()
	go recvResponse(ctx, clientStream)
	return done
	// <-done
}

func sendRequests(ctx context.Context, clientStream proto.AppService_GreetClient) {
	persons := []*proto.PersonName{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Kannan"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ganesh", LastName: "Kumar"},
	}

	// done := make(chan struct{})

	for _, person := range persons {
		req := &proto.GreetRequest{
			Person: person,
		}
		log.Printf("Sending Person : %s %s\n", person.FirstName, person.LastName)
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func recvResponse(ctx context.Context, clientStream proto.AppService_GreetClient) {
	for {
		res, err := clientStream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.GetMessage())
	}
}
