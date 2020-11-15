package api

import (
	"fmt"
	"github.com/JIakki/genesis/api/controllers"
	"github.com/JIakki/genesis/db"
	"github.com/JIakki/genesis/modules/payment"
	"net/http"
)

func NewPaymentHandler(mux *http.ServeMux, database *db.DB) *http.ServeMux {

	repo := payment.NewRepository(payment.NewMapper(database))
	mux.HandleFunc("/payment/buttons", func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.NewGetPaymentButtonsCtrl(repo)
		fmt.Println(ctrl.StartedAt)

		switch r.Method {
		case http.MethodPost:
			ctrl.Execute(w, r)
		default:
			ctrl.Execute(w, r)
		}
	})

	return mux
}
