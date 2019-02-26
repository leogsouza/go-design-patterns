package factory

import (
	"errors"
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
	return nil, errors.New("Not implemented yet")
}

// CashPM corresponds to Cash payment Method
type CashPM struct{}

// DebitCardPM corresponds to Debit Card payment method
type DebitCardPM struct{}

// Pay implements payment method for CashPM
func (c *CashPM) Pay(amount float32) string {
	return ""
}

// Pay implements payment method for DebitCardPM
func (c *DebitCardPM) Pay(amount float32) string {
	return ""
}
