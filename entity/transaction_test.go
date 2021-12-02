package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTRansactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountId = "1"
	transaction.Amount = 2000
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "You dont have limit for this transaction", err.Error())
}

func TestTRansactionWithAmountLesserThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountId = "1"
	transaction.Amount = 0
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "The amount must be greater than 1", err.Error())
}
