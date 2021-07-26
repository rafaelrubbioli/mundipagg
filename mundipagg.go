package mundipagg

import (
	"net/http"
	"time"
)

type Mundipagg interface {
	NewSubscription(subscription *Subscription, uuid string) (*Subscription, error)
	NewCustomer(customer *Customer, uuid string) (*Customer, error)
	NewCardByToken(customerID string, cardToken string, uuid string) (*CreditCard, error)
	UpdateStartAt(input *string, subscriptionID string, uuid string) error
	UpdateNextBillingDay(nextBillingDay *time.Time, customerID string, uuid string) error
	AddDiscount(billExtras *BillExtras, subscriptionID string, uuid string) error
}

type mundipagg struct {
	BasicSecretAuthKey string
}

func New(secret string) Mundipagg {
	return mundipagg{
		BasicSecretAuthKey: secret,
	}
}

func (m mundipagg) NewCardByToken(customerID string, cardToken string, uuid string) (*CreditCard, error) {
	card := &CreditCardToken{
		Token: cardToken,
	}

	resp, err := Do(http.MethodPost, card, m.BasicSecretAuthKey, uuid, BASEURL+customerID+"/cards")
	if err != nil {
		return nil, err
	}

	result, err := NewCreditCard(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m mundipagg) NewCustomer(c *Customer, uuid string) (*Customer, error) {
	resp, err := Do(http.MethodPost, c, m.BasicSecretAuthKey, uuid, BASEURL+"customers")
	if err != nil {
		return nil, err
	}

	result, err := NewCustomer(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m mundipagg) AddDiscount(extras *BillExtras, subscriptionID string, uuid string) error {
	_, err := Do(http.MethodPost, extras, m.BasicSecretAuthKey, uuid, BASEURL+"subscriptions/"+subscriptionID+"/discounts")
	return err
}

func (m mundipagg) NewSubscription(s *Subscription, uuid string) (*Subscription, error) {
	resp, err := Do(http.MethodPost, s, m.BasicSecretAuthKey, uuid, BASEURL+"subscriptions")
	if err != nil {
		return nil, err
	}

	result, err := NewSubscription(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m mundipagg) UpdateStartAt(startAt *string, subscriptionID string, uuid string) error {
	input := struct {
		StartAt *string `json:"start_at,omitempty"`
	}{
		StartAt: startAt,
	}

	_, err := Do(http.MethodPatch, input, m.BasicSecretAuthKey, uuid, BASEURL+"subscriptions/"+subscriptionID+"/start-at")
	return err
}

func (m mundipagg) UpdateNextBillingDay(nextBillingDay *time.Time, customerID string, indepotencyKey string) error {
	input := struct {
		NextBillingDay *time.Time `json:"next_billing_at,omitempty"`
	}{
		NextBillingDay: nextBillingDay,
	}

	_, err := Do(http.MethodPatch, input, m.BasicSecretAuthKey, indepotencyKey, BASEURL+"subscriptions/"+customerID+"/billing-date")
	return err
}
