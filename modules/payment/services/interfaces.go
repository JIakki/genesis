package services

import "sync"

type IPaymentService interface {
	GetButton(price int, wg *sync.WaitGroup, button chan GetButtonResponse, err chan error)
}
