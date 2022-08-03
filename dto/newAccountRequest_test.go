package dto

import (
	"capi/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountRequest_Validate(t *testing.T) {
	testCase := []struct {
		name  string
		input NewAccountRequest
		want  *errs.AppErr
	}{
		{
			"valid transaction",
			NewAccountRequest{
				Amount:      5000.01,
				AccountType: "saving",
			},
			nil,
		},
		{
			"invalid account type",
			NewAccountRequest{
				Amount:      5000.01,
				AccountType: "invalid type",
			},
			errs.NewValidationError(ErrInvalidAccountType),
		},
		{
			"invalid amount",
			NewAccountRequest{
				Amount:      500,
				AccountType: "saving",
			},
			errs.NewValidationError(ErrInvalidAmount),
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			// act
			res := tt.input.Validate()

			// assert
			assert.Equal(t, tt.want, res)
		})
	}
}
