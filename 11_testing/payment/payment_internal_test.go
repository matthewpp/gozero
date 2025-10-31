package payment

import (
	"testing"
)

func TestPaymentWithCreditCard(t *testing.T) {
	t.Run("not accept 001 card", func(t *testing.T) {
		// p := Payment{}
		// p.PayWithCreditCard(500, "00123456")
		err := paymentWithCreditCard(500, "00123456")
		if err == nil {
			t.Errorf("err must not nil")
		}
	})

	t.Run("not accept 003 card", func(t *testing.T) {
		err := paymentWithCreditCard(500, "00323456")
		if err == nil {
			t.Errorf("err must not nil")
		}
	})

	t.Run("not accept james bond card", func(t *testing.T) {
		err := paymentWithCreditCard(500, "00723456")
		if err == nil {
			t.Errorf("err must not nil")
		}
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
