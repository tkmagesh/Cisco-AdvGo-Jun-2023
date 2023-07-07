package services

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"log"
	"time"
)

// overriding the UnimplementedAppServiceServer.Add method
/*
func (asi *AppService) Add(ctx context.Context, req *proto.MathOperationRequest) (*proto.MathOperationResponse, error) {
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
*/

// handing timeouts
func (asi *AppService) Add(ctx context.Context, req *proto.MathOperationRequest) (*proto.MathOperationResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("[appService.add()] processing %d and %d\n", x, y)

	time.Sleep(5 * time.Second) // forcing the timeout signal to be triggerd from the client
	select {
	case <-ctx.Done():
		log.Println("timeout occurred")
		return nil, errors.New("timeout occurred")
	default:
		result := x + y
		res := &proto.MathOperationResponse{
			Result: result,
		}
		log.Println("sending response")
		return res, nil
	}

}
