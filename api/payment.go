package api

import (
	"github.com/JIakki/genesis/api/controllers"
	"github.com/JIakki/genesis/db"
	"github.com/JIakki/genesis/modules/payment"
	"net/http"
)

func NewPaymentHandler(mux *http.ServeMux, database *db.DB) {

	repo := payment.NewRepository(payment.NewMapper(database))
	ctrl := controllers.NewGetPaymentButtonsCtrl(repo)
	mux.HandleFunc("/payment/buttons", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ctrl.Execute(w, r)
		default:
			ctrl.NotFound(w, r, "Not Found")
		}
	})
}
