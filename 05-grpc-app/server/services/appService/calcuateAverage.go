package services

import (
	"grpc-app/proto"
	"io"
	"log"
)

func (asi *AppService) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum, count int32

	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			average := sum / count
			res := &proto.AverageResponse{
				Average: average,
			}
			if err := serverStream.SendAndClose(res); err != nil {
				log.Fatalln(err)
			}
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		no := req.GetNo()
		log.Printf("Received CalculateAverage request for no = %d\n", no)
		sum += req.GetNo()
		count++
	}
	return nil
}
