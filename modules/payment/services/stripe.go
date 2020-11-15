package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type StripeService struct {
	pubKey string
}

func (service *StripeService) GetButton(price int, wg *sync.WaitGroup, buttonChan chan GetButtonResponse, errorChan chan error) {
	defer wg.Done()
	var result map[string]interface{}
	resp, err := http.Get(fmt.Sprintf("https://gock.com/payment/stripe/buttons/%d?pubKey=%s", price, service.pubKey))
	if err != nil {
		errorChan <- err
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		errorChan <- err
		return
	}
	name, ok := result["name1"].(string)
	if !ok {
		errorChan <- fmt.Errorf("Stripe name does not exists")
		return
	}

	url, ok := result["url1"].(string)
	if !ok {
		errorChan <- fmt.Errorf("Stripe url does not exists")
		return
	}

	buttonChan <- NewButtonFormatter().Format(&PaymentButton{Name: name, URL: url})
}

func NewStripeService(pubKey string) *StripeService {
	return &StripeService{pubKey}
}
