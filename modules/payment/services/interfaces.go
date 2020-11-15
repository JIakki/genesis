package services

type IPaymentService interface {
	GetButton(price int) (*PaymentButton, error)
}
