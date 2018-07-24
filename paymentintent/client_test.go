package paymentintent

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestPaymentIntentCancel(t *testing.T) {
	intent, err := Cancel("pi_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentCapture(t *testing.T) {
	intent, err := Capture("pi_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentConfirm(t *testing.T) {
	intent, err := Confirm("pi_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentGet(t *testing.T) {
	intent, err := Get("pi_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentList(t *testing.T) {
	i := List(&stripe.PaymentIntentListParams{})

	// Verify that we can get at least one payment intent
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.PaymentIntent())
}

func TestPaymentIntentNew(t *testing.T) {
	intent, err := New(&stripe.PaymentIntentParams{
		AllowedSourceTypes: []*string{
			stripe.String("card"),
		},
		Amount:   stripe.Int64(123),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentUpdate(t *testing.T) {
	intent, err := Update("tr_123", &stripe.PaymentIntentParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}