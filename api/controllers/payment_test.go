package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JIakki/genesis/db"
	"github.com/JIakki/genesis/modules/payment"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPaymentButtons(t *testing.T) {
	database := db.New()
	repo := payment.NewRepository(payment.NewMapper(database))
	ctrl := NewGetPaymentButtonsCtrl(repo)

	t.Run("returns Pepper's score", func(t *testing.T) {
		var result []map[string]interface{}
		gock.New("https://gock.com").
			Get("/payment/buttons/1").
			Reply(200).
			JSON(map[string]string{"name": "paypal", "url": "paypal-url"})

		defer gock.Off()
		request, _ := http.NewRequest(http.MethodPost, "/payment/buttons", nil)
		response := httptest.NewRecorder()

		ctrl.Execute(response, request)
		json.Unmarshal([]byte(response.Body.String()), &result)
		fmt.Println(response.Body, "---")

		if response.Code != 200 {
			t.Errorf("%s", result)
		}

		if result[0]["url"] != "paypal-url" {
			t.Errorf("got %q, want %q", result[0]["url"], "paypal-url")
		}
	})

	t.Run("returns 500 if response from services is invalid", func(t *testing.T) {
		gock.New("https://gock.com").
			Get("/payment/buttons/1").
			Reply(200).
			JSON(map[string]string{"fname": "paypal", "url": "paypal-url"})
		defer gock.Off()
		request, _ := http.NewRequest(http.MethodPost, "/payment/buttons", nil)
		response := httptest.NewRecorder()

		ctrl.Execute(response, request)

		if response.Code != 500 {
			t.Errorf("Status should be 500")
		}
	})
}
