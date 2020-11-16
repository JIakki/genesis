package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PayPalService struct {
	pubKey string
	price  int
}

func (s *PayPalService) GetButton(ctx context.Context) (button GetButtonResponse, err error) {
	var result map[string]interface{}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://gock.com/payment/paypal/buttons/%d?pubKey=%s", s.price, s.pubKey), nil)
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
	name, ok := result["name"].(string)
	if !ok {
		return button, fmt.Errorf("Paypal name does not exists")
	}

	url, ok := result["url"].(string)
	if !ok {
		return button, fmt.Errorf("Paypal url does not exists")
	}

	return NewButtonFormatter().Format(&PaymentButton{Name: name, URL: url}), nil
}

func NewPayPalService(pubKey string, price int) *PayPalService {
	return &PayPalService{pubKey, price}
}
