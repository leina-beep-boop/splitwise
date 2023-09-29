package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	lpb "github.com/leina-beep-boop/splitwise/protos/ledger"
	spb "github.com/leina-beep-boop/splitwise/protos/splitter"
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
	sc := spb.NewSplitterServiceClient(conn)

	req := lpb.LedgerRequest{
		Expense: &lpb.Expense{
			PersonId:    "1",
			Debtors:     []string{"2", "3", "4"},
			ExpenseBool: true,
			Amount:      12.39,
			Description: "test expense",
		},
	}
	// communicate to server
	lresp, err := lc.AddExpense(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling AddExpense: %s", err)
	}
	log.Printf("Received response message from server: %s", lresp.Status)

	sresp, err := sc.CalculateDebt(context.Background(), &spb.SplitterRequest{})
	if err != nil {
		log.Fatalf("Error when calling AddExpense: %s", err)
	}
	log.Printf("Received response message from server: %s", sresp.Status)
}
