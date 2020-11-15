package controllers

import (
	"github.com/JIakki/genesis/core"
	"github.com/JIakki/genesis/modules/payment"
	"github.com/JIakki/genesis/modules/payment/services"
	"net/http"
)

type GetPaymentButtonsCtrl struct {
	ProductRepo payment.IProductRepository
	*core.BaseController
}

func (ctrl *GetPaymentButtonsCtrl) Execute(w http.ResponseWriter, r *http.Request) {
	p, err := ctrl.ProductRepo.FindById(1)
	if err != nil {
		ctrl.InternalError(w, r, err)
	}

	aggregator := services.NewAggregator([]services.IPaymentService{
		services.NewPayPalService("Key"),
	})

	buttons, err := aggregator.Aggregate(p.Price)
	if err != nil {
		ctrl.InternalError(w, r, err)
	}

	ctrl.JSON(w, r, buttons)
}

func NewGetPaymentButtonsCtrl(repo payment.IProductRepository) *GetPaymentButtonsCtrl {
	return &GetPaymentButtonsCtrl{
		ProductRepo:    repo,
		BaseController: core.NewBaseController(),
	}
}
