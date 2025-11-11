package client

// go directive
//
//go:generate go run go.uber.org/mock/mockgen -source=./client.go -destination=./mock_client/client.go
type Client interface {
	GetBalance(cardNumber string) (int, error)
}

type client struct {
}

func (c client) GetBalance(cardNumber string) (int, error) {
	// Do something to get balance from external service

	return 0, nil
}

func New() Client {
	return &client{}
}
