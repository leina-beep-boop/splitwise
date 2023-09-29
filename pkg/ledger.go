package pkg

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	pb "github.com/leina-beep-boop/splitwise/protos/ledger"
)

const dataPath string = "./data/expenses.csv"

type LedgerServiceServer struct {
	pb.UnimplementedLedgerServiceServer
}

func (LedgerServiceServer) AddExpense(ctx context.Context, req *pb.LedgerRequest) (*pb.LedgerResponse, error) {
	file, err := os.OpenFile(dataPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Could not open dataset at path: %v.\nErr: %v", dataPath, err)
		return &pb.LedgerResponse{Status: 2}, err
	}
	exp := req.Expense
	file.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v\n",
		exp.Id,
		exp.PersonId,
		exp.ExpenseBool,
		exp.Amount,
		exp.Description))
	log.Printf("add expense %v", exp)
	file.Close()
	return &pb.LedgerResponse{Status: 1}, nil
}

func (LedgerServiceServer) ResetExpenses(ctx context.Context, req *pb.LedgerRequest) (*pb.LedgerResponse, error) {
	file, err := os.Create(dataPath)
	if err != nil {
		log.Printf("Could not create dataset at path: %v.\nErr: %v", dataPath, err)
		return &pb.LedgerResponse{Status: 2}, err
	}
	file.WriteString("id,person_id,expense_bool,amount,description\n")
	file.Close()
	return &pb.LedgerResponse{Status: 1}, nil
}

func (LedgerServiceServer) GetAllExpenses(ctx context.Context, req *pb.LedgerRequest) (*pb.LedgerResponse, error) {
	file, err := os.Open(dataPath)
	if err != nil {
		log.Printf("Could not open dataset at path: %v.\nErr: %v", dataPath, err)
		return &pb.LedgerResponse{Status: 2}, err
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Printf("Could not read dataset at path: %v.\nErr: %v", dataPath, err)
		return &pb.LedgerResponse{Status: 2}, err
	}
	expenses := make([]*pb.Expense, len(records)-1) // first row is header
	var exp pb.Expense
	for i, record := range records {
		if i != 0 {
			amt, _ := strconv.ParseFloat(record[3], 64)
			expense_bool, _ := strconv.ParseBool(record[2])
			exp = pb.Expense{
				Id:          record[0],
				PersonId:    record[1],
				ExpenseBool: expense_bool,
				Amount:      amt,
				Description: record[4],
			}
			expenses[i-1] = &exp // account for header
		}
	}
	return &pb.LedgerResponse{Expenses: expenses, Status: 1}, nil
}
