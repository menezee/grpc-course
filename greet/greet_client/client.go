package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-course/greet/greetpb"
	"log"
)

func main() {
	fmt.Println("Hello, I'm a client")

	// WithInsecure because locally we are not using ssl
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer clientConnection.Close()

	c := greetpb.NewGreetServiceClient(clientConnection)
	fmt.Printf("Created client: %f", c)
}
