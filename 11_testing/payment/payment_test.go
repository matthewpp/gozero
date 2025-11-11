package payment_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"gozero/chonlatee/payment"
	"gozero/chonlatee/payment/client/mock_client"
)

func TestPayment_PayWithCreditCard(t *testing.T) {
	t.Run("balance not enough", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// dependency injection
		balanceClient := mock_client.NewMockClient(ctrl)
		p := payment.Payment{
			BalanceClient: balanceClient,
		}

		cardNumber := "00123456"
		balanceClient.EXPECT().GetBalance(cardNumber).Return(5, nil)

		err := p.PayWithCreditCard(cardNumber)

		assert.Error(t, err)
		assert.Equal(t, "your balance is not enough", err.Error())
	})

	t.Run("not accept 001 card", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// dependency injection
		balanceClient := mock_client.NewMockClient(ctrl)
		p := payment.Payment{
			BalanceClient: balanceClient,
		}

		cardNumber := "00123456"
		balanceClient.EXPECT().GetBalance(cardNumber).Return(100, nil)

		err := p.PayWithCreditCard(cardNumber)

		assert.Error(t, err)
		assert.Equal(t, "we not accept card start with 001", err.Error())
	})

	t.Run("not accept 003 card", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// dependency injection
		balanceClient := mock_client.NewMockClient(ctrl)
		p := payment.Payment{
			BalanceClient: balanceClient,
		}

		cardNumber := "00323456"
		balanceClient.EXPECT().GetBalance(cardNumber).Return(100, nil)

		err := p.PayWithCreditCard(cardNumber)

		assert.Error(t, err)
		assert.Equal(t, "we not accept card start with 003", err.Error())
	})

	t.Run("not accept james bond card", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// dependency injection
		balanceClient := mock_client.NewMockClient(ctrl)
		p := payment.Payment{
			BalanceClient: balanceClient,
		}

		cardNumber := "00723456"
		balanceClient.EXPECT().GetBalance(cardNumber).Return(100, nil)

		err := p.PayWithCreditCard(cardNumber)

		assert.Error(t, err)
		assert.Equal(t, "we not accept card from james bond", err.Error())
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			p := &payment.Payment{
				BalanceClient: mock_client.NewMockClient(ctrl),
			}
			err := p.Pay(tc.price)
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("expect %v but got %v", tc.expectErr, err)
			}
		})
	}

}
