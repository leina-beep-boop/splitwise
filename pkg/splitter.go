package pkg

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	pb "github.com/leina-beep-boop/splitwise/protos/splitter"
)

type SplitterServiceServer struct {
	pb.UnimplementedSplitterServiceServer
}

// TODO: try implementing with proto.Marshall and UnMarshall
func (SplitterServiceServer) CalculateDebt(ctx context.Context, req *pb.SplitterRequest) (*pb.SplitterResponse, error) {
	log.Printf("Hello World! %v", req)

	file, err := os.Open(dataPath)
	if err != nil {
		log.Printf("Could not open dataset at path: %v.\nErr: %v", dataPath, err)
		return &pb.SplitterResponse{Status: 2}, err
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Printf("Could not read dataset at path: %v.\nErr: %v", dataPath, err)
		return &pb.SplitterResponse{Status: 2}, err
	}
	accounts, _ := createAccounts(records)
	for i, a := range accounts {
		log.Printf("Account %v: %v", i, a)
	}

	for _, r := range records {
		AddDebt(accounts, r)
	}
	return &pb.SplitterResponse{Status: 0}, nil
}

func createAccounts(records [][]string) ([]*pb.Account, error) {
	accounts := make([]*pb.Account, 0, len(records))
	unique_accounts := make(map[string]bool)

	for _, r := range records[1:] {
		if _, unique := unique_accounts[r[1]]; !unique {
			unique_accounts[r[1]] = true
			a := pb.Account{
				PersonId: r[1],
				Debts:    make([]*pb.Debt, 0, len(records)),
			}
			accounts = append(accounts, &a)
		}
		// check debtors
		if _, unique := unique_accounts[r[2]]; !unique {
			unique_accounts[r[2]] = true
			a := pb.Account{
				PersonId: r[2],
				Debts:    make([]*pb.Debt, 0, len(records)),
			}
			accounts = append(accounts, &a)
		}
	}
	return accounts, nil
}

func AddDebt(a []*pb.Account, record []string) ([]*pb.Account, error) {
	person_id := record[1]
	debtor_id := record[2]
	expense_bool, _ := strconv.ParseBool(record[3])
	amount, _ := strconv.ParseFloat(record[4], 64)
	if !expense_bool {
		amount = amount * -1
	}

	return a, nil
}
