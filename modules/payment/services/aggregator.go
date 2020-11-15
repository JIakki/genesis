package services

type Aggregator struct {
	services []IPaymentService
}

func (aggregator *Aggregator) Aggregate(price int) ([]*PaymentButton, error) {
	var res []*PaymentButton
	for _, service := range aggregator.services {
		button, err := service.GetButton(price)
		if err != nil {
			return nil, err
		}

		res = append(res, button)
	}
	return res, nil
}

func NewAggregator(services []IPaymentService) *Aggregator {
	return &Aggregator{services}
}
