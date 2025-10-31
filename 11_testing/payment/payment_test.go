package payment_test

import (
	"gozero/chonlatee/payment"
	"testing"
)

func TestPayment_PayWithCreditCard(t *testing.T) {
	t.Run("not accept 001 card", func(t *testing.T) {
		p := payment.Payment{}
		err := p.PayWithCreditCard(500, "00123456")
		if err == nil {
			t.Errorf("err must not nil")
		}
	})

	t.Run("not accept 003 card", func(t *testing.T) {
		p := payment.Payment{}
		err := p.PayWithCreditCard(500, "00323456")
		if err == nil {
			t.Errorf("err must not nil")
		}
	})

	t.Run("not accept james bond card", func(t *testing.T) {
		p := payment.Payment{}
		err := p.PayWithCreditCard(500, "00723456")
		if err == nil {
			t.Errorf("err must not nil")
		}
	})
}

func TestPayment_Pay(t *testing.T) {

	testCases := []struct {
		name      string
		price     int
		expectErr error
	}{
		{
			name:      "error lower pric",
			price:     99,
			expectErr: payment.ErrPayLower,
		},
		{
			name:      "error over pric",
			price:     5000,
			expectErr: payment.ErrPayOver,
		},
		{
			name:      "pay success",
			price:     500,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := &payment.Payment{}
			err := p.Pay(tc.price)
			if err != tc.expectErr {
				t.Errorf("expect %v but got %v", tc.expectErr, err)
			}
		})
	}

}
