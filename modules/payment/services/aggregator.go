package services

import (
	"context"
	"sync"
)

type Aggregator struct {
	services []IPaymentService
}

func (aggregator *Aggregator) Aggregate(ctx context.Context) ([]GetButtonResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	var res []GetButtonResponse

	done := make(chan struct{})
	errChan := make(chan error, len(aggregator.services))
	mu := sync.Mutex{}

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	for _, service := range aggregator.services {
		wg.Add(1)
		go func(service IPaymentService) {
			defer wg.Done()

			button, err := service.GetButton(ctx)
			if err != nil {
				errChan <- err
				return
			}
			mu.Lock()
			res = append(res, button)
			mu.Unlock()
		}(service)
	}

	select {
	case err := <-errChan:
		return nil, err
	case <-done:
		return res, nil
	}
}

func NewAggregator(services []IPaymentService) *Aggregator {
	return &Aggregator{services}
}
