package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StripeService struct {
	pubKey string
	price  int
}

func (s *StripeService) GetButton(ctx context.Context) (button GetButtonResponse, err error) {
	var result map[string]interface{}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://gock.com/payment/stripe/buttons/%d?pubKey=%s", s.price, s.pubKey), nil)
	if err != nil {
		return button, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return button, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return button, err
	}

	name, ok := result["name1"].(string)
	if !ok {
		return button, fmt.Errorf("Stripe name does not exists")
	}

	url, ok := result["url1"].(string)
	if !ok {
		return button, fmt.Errorf("Stripe url does not exists")
	}

	return NewButtonFormatter().Format(&PaymentButton{Name: name, URL: url}), nil
}

func NewStripeService(pubKey string, price int) *StripeService {
	return &StripeService{pubKey, price}
}
