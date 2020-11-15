package services

import "sync"

type IPaymentService interface {
	GetButton(price int, wg *sync.WaitGroup, button chan *PaymentButton, err chan error)
}
