package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PayPalService struct {
	pubKey string
}

func (service *PayPalService) GetButton(price int) (*PaymentButton, error) {
	var result map[string]interface{}
	resp, err := http.Get("https://gock.com/payment/buttons/1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, err
	}
	name, ok := result["name"].(string)
	if !ok {
		return nil, fmt.Errorf("Paypay name does not exists")
	}

	url, ok := result["url"].(string)
	if !ok {
		return nil, fmt.Errorf("Paypay url does not exists")
	}

	return &PaymentButton{Name: name, URL: url}, err
}

func NewPayPalService(pubKey string) *PayPalService {
	return &PayPalService{pubKey}
}
