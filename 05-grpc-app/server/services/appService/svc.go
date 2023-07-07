package services

import (
	"grpc-app/proto"
)

type AppService struct {
	proto.UnimplementedAppServiceServer
}
