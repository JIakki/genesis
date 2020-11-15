package api

import (
	"github.com/JIakki/genesis/api/controllers"
	"github.com/JIakki/genesis/db"
	"github.com/JIakki/genesis/modules/payment"
	"net/http"
)

func NewPaymentHandler(mux *http.ServeMux, database *db.DB) *http.ServeMux {

	repo := payment.NewRepository(payment.NewMapper(database))
	mux.HandleFunc("/payment/buttons", func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.NewGetPaymentButtonsCtrl(repo)

		switch r.Method {
		case http.MethodPost:
			ctrl.Execute(w, r)
		default:
			ctrl.NotFound(w, r, "404")
		}
	})

	mux.HandleFunc("/payment/callback/{service}", func(w http.ResponseWriter, r *http.Request) {
		// UseAuthMiddleWare(w, r)
		// controllers.NewPaymentCallbackCtrl(repo)
		// switch r.Method {
		// case http.MethodPost:
		// ctrl.Execute(w, r)
		// default:
		// ctrl.NotFound(w, r, "404")
		// }
	})

	return mux
}
