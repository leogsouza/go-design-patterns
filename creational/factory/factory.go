package factory

import (
	"errors"
	"fmt"
)

// PaymentMethod is a interface that wraps a method Pay
type PaymentMethod interface {
	Pay(amount float32) string
}

// Payment method types
const (
	Cash      = 1
	DebitCard = 2
)

// GetPaymentMethod returns the PaymentMethod or an error
func GetPaymentMethod(m int) (PaymentMethod, error) {

	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n",
			m))
	}
}

// CashPM corresponds to Cash payment Method
type CashPM struct{}

// DebitCardPM corresponds to Debit Card payment method
type DebitCardPM struct{}

// Pay implements payment method for CashPM
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// Pay implements payment method for DebitCardPM
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}
