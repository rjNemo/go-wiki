package service

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

// secret API key
var secretKey string = "sk_test_dyp7eLEq2KnxpWE0qIMihTYZ"

// service.PaymentIntent(1000, "ruidy.nemausat@gmail.com")

// PaymentIntent creates a	payment intent
func PaymentIntent(a int64, m string) (*stripe.PaymentIntent, error) {
	stripe.Key = secretKey
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(a),
		Currency:           stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		ReceiptEmail:       stripe.String(m),
	}
	return paymentintent.New(params)
}
