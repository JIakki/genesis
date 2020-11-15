package controllers

import (
	// "fmt"
	"github.com/JIakki/genesis/core"
	"github.com/JIakki/genesis/modules/payment"
	"net/http"
)

type GetPaymentButtonsCtrl struct {
	ProductRepo payment.IProductRepository
	core.BaseController
}

func (ctrl *GetPaymentButtonsCtrl) Execute(w http.ResponseWriter, r *http.Request) {
	p, err := ctrl.ProductRepo.FindById(1)
	if err != nil {
		ctrl.InternalError(w, r, p)
	}

	ctrl.JSON(w, r, p)
}

func NewGetPaymentButtonsCtrl(repo payment.IProductRepository) *GetPaymentButtonsCtrl {
	return &GetPaymentButtonsCtrl{
		ProductRepo: repo,
	}
}
