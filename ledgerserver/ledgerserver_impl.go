package ledgerserver

import (
	"context"
	"log"

	"github.com/leina-beep-boop/splitwise/ledger"
	pb "github.com/leina-beep-boop/splitwise/protos"
)

type LedgerServiceServer struct {
	pb.UnimplementedLedgerServiceServer
}

func New() pb.LedgerServiceServer {
	return &LedgerServiceServer{}
}

func (*LedgerServiceServer) AddExpense(ctx context.Context, req *pb.LedgerRequest) (*pb.LedgerResponse, error) {
	err := ledger.New().AddExpense(req.Expense)
	if err != nil {
		return &pb.LedgerResponse{Status: 2}, err
	}
	log.Printf("add expense %v", req.Expense)
	return &pb.LedgerResponse{Status: 1}, nil
}

func (*LedgerServiceServer) ResetExpenses(ctx context.Context, req *pb.LedgerRequest) (*pb.LedgerResponse, error) {
	err := ledger.New().ResetExpenses()
	if err != nil {
		return &pb.LedgerResponse{Status: 2}, err
	}
	log.Printf("expenses reset")
	return &pb.LedgerResponse{Status: 1}, nil
}

func (*LedgerServiceServer) GetAllExpenses(ctx context.Context, req *pb.LedgerRequest) (*pb.LedgerResponse, error) {
	expenses, err := ledger.New().ReadExpenses()
	if err != nil {
		return &pb.LedgerResponse{Status: 2}, err
	}
	return &pb.LedgerResponse{Expenses: expenses, Status: 1}, nil
}
