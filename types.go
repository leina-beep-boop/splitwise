package main

import "context"

type Expense struct {
	id           int32   `json:"id"`
	person_id    int32   `json:"person_id"`
	expense_bool bool    `json:"expense_bool"`
	amount       float32 `json:"amount"`
	description  string  `'json:"string"`
}

type AllExpenses struct {
	expenses []Expense `json:"expenses"`
}

type LedgerRequest struct {
	Expense Expense `json:"expense"`
}

type LedgerResponse struct {
	expenses AllExpenses `json:"expenses"`
	status   int         `json:"status"`
}

type LedgerService interface {
	AddExpense(context.Context, *LedgerRequest) (*LedgerResponse, error)
	ResetExpense(context.Context, *LedgerRequest) (*LedgerResponse, error)
	GetAllExpenses(context.Context, *LedgerRequest) (*LedgerResponse, error)
}
