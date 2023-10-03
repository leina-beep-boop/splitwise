package ledger

import (
	"encoding/json"
	"log"
	"os"

	pb "github.com/leina-beep-boop/splitwise/protos"
)

const dataPath string = "./data/expenses.json"

type Ledger struct {
	expenses []*pb.Expense
}

// enable optional arguments on Ledger
// see this for more details: https://charlesxu.io/go-opts/
type Option func(led *Ledger)

func WithExpenses(exps []*pb.Expense) Option {
	return func(led *Ledger) {
		led.expenses = exps
	}
}

func New(opts ...Option) *Ledger {
	led := &Ledger{}
	for _, opt := range opts {
		opt(led)
	}
	return led
}

func (Ledger) CalculateDebt() error {
	return nil
}

func (led *Ledger) ReadExpenses() ([]*pb.Expense, error) {
	file, err := os.ReadFile(dataPath)
	if err != nil {
		log.Printf("Could not open dataset at path: %v.\nErr: %v", dataPath, err)
	}
	expenses := []*pb.Expense{}
	err = json.Unmarshal(file, &expenses)
	if err != nil {
		log.Printf("Could not unmarshal dataset. Err: %v", err)
	}
	return expenses, err
}

func (led *Ledger) AddExpense(exp *pb.Expense) error {
	// This implementation is O(n) and should be O(1). This obviously would
	// not scale well. When the data gets there, switch to an actual database.

	// append existing data to new request data
	exps, _ := led.ReadExpenses()
	exps = append(exps, exp)
	// write to json dataset
	expBytes, err := json.MarshalIndent(exps, " ", " ")
	if err != nil {
		log.Printf("Could not marshall json expenses: %v.\nErr: %v", led.expenses, err)
	}
	err = os.WriteFile(dataPath, expBytes, 0644)
	if err != nil {
		log.Printf("Could not write dataset at path: %v.\nErr: %v", dataPath, err)
	}
	return nil
}

func (*Ledger) ResetExpenses() error {
	_, err := os.Create(dataPath)
	if err != nil {
		log.Printf("Could not create dataset at path: %v.\nErr: %v", dataPath, err)
		return err
	}
	return nil
}
