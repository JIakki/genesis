package api

import (
	"net/http"
	"fmt"
)

func NewPaymentHandler(mux *http.ServeMux) {
	mux.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "data")
	})
}
