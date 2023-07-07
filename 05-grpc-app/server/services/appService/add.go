package services

import (
	"context"
	"fmt"
	"grpc-app/proto"
)

// overriding the UnimplementedAppServiceServer.Add method
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
