package services

import (
	"context"
)

type IPaymentService interface {
	GetButton(ctx context.Context) (GetButtonResponse, error)
}
