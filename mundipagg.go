package mundipagg

import (
	"net/http"
	"time"
)

type Mundipagg interface {
	NewSubscription(subscription *Subscription, uuid string) (*Response, error)
	NewCustomer(customer *Customer, uuid string) (*Response, error)
	NewCardByToken(customerID string, cardToken string, uuid string) (*Response, error)
	UpdateStartAt(input *string, subscriptionID string, uuid string) (*Response, error)
	UpdateNextBillingDay(nextBillingDay *time.Time, customerID string, uuid string) (*Response, error)
	AddDiscount(billExtras *BillExtras, subscriptionID string, uuid string) (*Response, error)
}

type mundipagg struct {
	BasicSecretAuthKey string
}

func New(secret string) Mundipagg {
	return mundipagg{
		BasicSecretAuthKey: secret,
	}
}

func (m mundipagg) NewCardByToken(customerID string, cardToken string, uuid string) (*Response, error) {
	card := &CreditCardToken{
		Token: cardToken,
	}

	resp, err := Do(http.MethodPost, card, m.BasicSecretAuthKey, uuid, BASEURL+customerID+"/cards")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m mundipagg) NewCustomer(c *Customer, uuid string) (*Response, error) {
	resp, err := Do(http.MethodPost, c, m.BasicSecretAuthKey, uuid, BASEURL+"customers")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m mundipagg) AddDiscount(extras *BillExtras, subscriptionID string, uuid string) (*Response, error) {
	resp, err := Do(http.MethodPost, extras, m.BasicSecretAuthKey, uuid, BASEURL+"subscriptions/"+subscriptionID+"/discounts")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m mundipagg) NewSubscription(s *Subscription, uuid string) (*Response, error) {
	resp, err := Do(http.MethodPost, s, m.BasicSecretAuthKey, uuid, BASEURL+"subscriptions")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m mundipagg) UpdateStartAt(startAt *string, subscriptionID string, uuid string) (*Response, error) {
	input := struct {
		StartAt *string `json:"start_at,omitempty"`
	}{
		StartAt: startAt,
	}

	resp, err := Do(http.MethodPatch, input, m.BasicSecretAuthKey, uuid, BASEURL+"subscriptions/"+subscriptionID+"/start-at")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m mundipagg) UpdateNextBillingDay(nextBillingDay *time.Time, customerID string, indepotencyKey string) (*Response, error) {
	input := struct {
		NextBillingDay *time.Time `json:"next_billing_at,omitempty"`
	}{
		NextBillingDay: nextBillingDay,
	}

	resp, err := Do(http.MethodPatch, input, m.BasicSecretAuthKey, indepotencyKey, BASEURL+"subscriptions/"+customerID+"/billing-date")
	if err != nil {
		return nil, err
	}

	return resp, nil
}
