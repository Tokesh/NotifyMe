package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "project/your/go/package" // Replace with the path to your proto package
)

const (
	address = "localhost:50051"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewYourServiceClient(conn)

	// Define the login credentials for the test
	login := "YourTestLogin"       // Replace with an actual login
	password := "YourTestPassword" // Replace with an actual password

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	log.Printf("Server listening at %v", address)
	// Call the Login RPC method
	resp, err := client.Login(ctx, &pb.YourLoginRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not login: %v", err)
	}

	log.Printf("Received token: %s", resp)
}
