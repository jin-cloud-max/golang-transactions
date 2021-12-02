package entity

import "errors"

type Transaction struct {
	ID           string
	AccountId    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}

}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("You dont have limit for this transaction")
	}

	if t.Amount < 1 {
		return errors.New("The amount must be greater than 1")
	}

	return nil
}
