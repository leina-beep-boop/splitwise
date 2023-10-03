package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	lpb "github.com/leina-beep-boop/splitwise/protos"
)

// TODO: add command line features to make it a CLI application

func main() {
	// establish stubby connection
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %s", err)
	}
	defer conn.Close()

	lc := lpb.NewLedgerServiceClient(conn)

	resp, _ := lc.ResetExpenses(context.Background(), &lpb.LedgerRequest{})
	log.Printf("Received response message from server: %s", resp.Status)

	req := lpb.LedgerRequest{
		Expense: &lpb.Expense{
			PersonId:    "1",
			Amount:      15.0,
			Description: "test expense",
		},
	}
	resp, err = lc.AddExpense(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling AddExpense: %s", err)
	}
	log.Printf("Received response message from server: %s", resp.Status)

	req = lpb.LedgerRequest{
		Expense: &lpb.Expense{
			PersonId:    "2",
			Amount:      5.0,
			Description: "test expense 2",
		},
	}
	resp, err = lc.AddExpense(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling AddExpense: %s", err)
	}
	log.Printf("Received response message from server: %s", resp.Status)

	resp, err = lc.GetAllExpenses(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling GetAllExpenses %s", err)
	}
	log.Printf("Received response message from server: %s", resp.Status)
	log.Printf("expenses: %s", resp.Expenses)
}
