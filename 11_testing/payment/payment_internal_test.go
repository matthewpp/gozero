package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentWithCreditCard(t *testing.T) {
	t.Run("not accept 001 card", func(t *testing.T) {
		err := paymentWithCreditCard("00123456")
		if err == nil {
			t.Errorf("err must not nil")
		}

		// Expected
		assert.Error(t, err)
		assert.Equal(t, "we not accept card start with 001", err.Error())
	})

	t.Run("not accept 003 card", func(t *testing.T) {
		err := paymentWithCreditCard("00323456")
		if err == nil {
			t.Errorf("err must not nil")
		}

		// Expected
		assert.Error(t, err)
		assert.Equal(t, "we not accept card start with 003", err.Error())
	})

	t.Run("not accept james bond card", func(t *testing.T) {
		err := paymentWithCreditCard("00723456")
		if err == nil {
			t.Errorf("err must not nil")
		}

		// Expected
		assert.Error(t, err)
		assert.Equal(t, "we not accept card from james bond", err.Error())
	})
}

func TestPay(t *testing.T) {
	testCases := []struct {
		name      string
		price     int
		expectErr error
	}{
		{
			name:      "error lower pric",
			price:     99,
			expectErr: ErrPayLower,
		},
		{
			name:      "error over pric",
			price:     5000,
			expectErr: ErrPayOver,
		},
		{
			name:      "pay success",
			price:     500,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := pay(tc.price)
			if err != tc.expectErr {
				t.Errorf("expect %v but got %v", tc.expectErr, err)
			}
		})
	}

}
