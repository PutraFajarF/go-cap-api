package dto

import (
	"capi/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionValidation_valid_transaction(t *testing.T) {
	// arrange
	txReq := TransactionRequest{
		TransactionType: WITHDRAWAL,
		Amount:          0,
	}
	// act
	res := txReq.Validate()

	// assert
	if res != nil {
		t.Error("error occured")
	}
}
func TestTransactionValidation_invalid_transaction_type(t *testing.T) {
	// arrange
	txReq := TransactionRequest{
		TransactionType: "invalid_type",
		Amount:          -1,
	}
	want := errs.NewValidationError("")
	// act
	res := txReq.Validate()

	// assert
	assert.NotNil(t, res)
	assert.Equal(t, want.Code, res.Code)

}

func TestTransactionValidation_invalid_transaction_amount(t *testing.T) {
	// arrange
	txReq := TransactionRequest{
		Amount:          -1,
		TransactionType: WITHDRAWAL,
	}
	want := errs.NewValidationError("")
	// act
	res := txReq.Validate()

	// assert
	assert.NotNil(t, res)
	assert.Equal(t, want.Code, res.Code)

}
