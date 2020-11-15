package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type PayPalService struct {
	pubKey string
}

func (service *PayPalService) GetButton(price int, wg *sync.WaitGroup, buttonChan chan *PaymentButton, errorChan chan error) {
	defer wg.Done()
	var result map[string]interface{}
	resp, err := http.Get(fmt.Sprintf("https://gock.com/payment/paypal/buttons/%d?pubKey=%s", price, service.pubKey))
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
	name, ok := result["name"].(string)
	if !ok {
		errorChan <- fmt.Errorf("Paypal name does not exists")
		return
	}

	url, ok := result["url"].(string)
	if !ok {
		errorChan <- fmt.Errorf("Paypal url does not exists")
		return
	}

	buttonChan <- &PaymentButton{Name: name, URL: url}
}

func NewPayPalService(pubKey string) *PayPalService {
	return &PayPalService{pubKey: pubKey}
}
