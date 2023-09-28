package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/leina-beep-boop/splitwise/protos/ledger"
)

func main() {
	// establish stubby connection
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewLedgerServiceClient(conn)

	// communicate to server
	resp, err := c.AddExpense(context.Background(), &pb.LedgerRequest{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Received response message from server: %s", resp.Body)
}
