package services

import (
	"fmt"
	"grpc-app/proto"
	"log"
	"time"
)

func (asi *AppService) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		greetReq, err := serverStream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		person := greetReq.GetPerson()
		firstName := person.GetFirstName()
		lastName := person.GetLastName()
		log.Printf("Received greet request for %q and %q\n", firstName, lastName)
		message := fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
		time.Sleep(2 * time.Second)
		log.Printf("Sending response : %q\n", message)
		greetResp := &proto.GreetResponse{
			Message: message,
		}
		if err := serverStream.Send(greetResp); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}
