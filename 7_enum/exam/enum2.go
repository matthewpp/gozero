package exam

type PaymentType string

// Using iota to define enum values
const (
	CreditCard   PaymentType = "CREDIT_CARD"
	PayPal       PaymentType = "PAYPAL"
	BankTransfer PaymentType = "BANK_TRANSFER"
)

func ToString(pt PaymentType) string {
	return string(pt)
}
