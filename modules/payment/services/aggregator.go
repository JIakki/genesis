package services

import (
	"sync"
)

type Aggregator struct {
	services []IPaymentService
}

func (aggregator *Aggregator) Aggregate(price int) ([]GetButtonResponse, error) {
	var wg sync.WaitGroup
	var res []GetButtonResponse

	done := make(chan bool)
	buttonChan := make(chan GetButtonResponse)
	errorChan := make(chan error)

	wg.Add(len(aggregator.services))
	go func() {
		wg.Wait()
		done <- true
	}()

	for _, service := range aggregator.services {
		go service.GetButton(price, &wg, buttonChan, errorChan)
	}

	for {
		select {
		case button := <-buttonChan:
			res = append(res, button)
		case err := <-errorChan:
			return nil, err
		case <-done:
			return res, nil
		}
	}
}

func NewAggregator(services []IPaymentService) *Aggregator {
	return &Aggregator{services}
}
