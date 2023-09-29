package pkg

import "context"

// this is actually useless, I got baited into writing this
// could be useful if I create another api that interacts
// with both grpc services
type Expense struct {
	Id           int32   `json:"id"`
	Person_id    int32   `json:"person_id"`
	Expense_bool bool    `json:"expense_bool"`
	Amount       float32 `json:"amount"`
	Description  string  `'json:"string"`
}

type AllExpenses struct {
	Expenses []Expense `json:"expenses"`
}

type LedgerRequest struct {
	Expense *Expense `json:"expense"`
}

type LedgerResponse struct {
	Expenses AllExpenses `json:"expenses"`
	Status   int         `json:"status"`
}

type LedgerService interface {
	AddExpense(context.Context, *LedgerRequest) (*LedgerResponse, error)
	ResetExpense(context.Context, *LedgerRequest) (*LedgerResponse, error)
	GetAllExpenses(context.Context, *LedgerRequest) (*LedgerResponse, error)
}
