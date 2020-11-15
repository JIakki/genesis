package controllers

import (
	"github.com/JIakki/genesis/db"
	"github.com/JIakki/genesis/modules/payment"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPaymentButtons(t *testing.T) {
	database := db.New()
	repo := payment.NewRepository(payment.NewMapper(database))
	ctrl := NewGetPaymentButtonsCtrl(repo)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/payment/buttons", nil)
		response := httptest.NewRecorder()

		ctrl.Execute(response, request)

		got := response.Body.String()
		want := "data"

		if response.Code != 200 {
			t.Errorf(got)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
