package api

import (
	"encoding/json"
	"github.com/JIakki/genesis/api/controllers"
	"github.com/JIakki/genesis/db"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPaymentButtons(t *testing.T) {
	database := db.New()
	mux := http.NewServeMux()
	router := NewPaymentHandler(mux, database)

	t.Run("returns buttons", func(t *testing.T) {
		var result controllers.GetButtonsResponse
		gock.New("https://gock.com").
			Get("/payment/paypal/buttons/100").
			Reply(200).
			JSON(map[string]string{"name": "paypal", "url": "paypal-url"})
		gock.New("https://gock.com").
			Get("/payment/stripe/buttons/100").
			Reply(200).
			JSON(map[string]string{"name1": "stripe", "url1": "stripe-url"})
		defer gock.Off()

		request, _ := http.NewRequest(http.MethodPost, "/payment/buttons", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		json.Unmarshal([]byte(writer.Body.String()), &result)

		if writer.Code != 200 {
			t.Errorf("%s", result)
		}
		url := result.Buttons[0].Url

		if url != "stripe-url" {
			t.Errorf("got %s, want %q", url, "stripe-url")
		}

		url = result.Buttons[1].Url
		if url != "paypal-url" {
			t.Errorf("got %s, want %q", url, "paypal-url")
		}

		if result.AppLink != "awasome-link-of-awasome-app" {
			t.Errorf("got %s, want %q", result.AppLink, "awasome-link-of-awasome-app")
		}

	})

	t.Run("returns 500 if writer from services is invalid", func(t *testing.T) {
		gock.New("https://gock.com").
			Get("/payment/paypal/buttons/100").
			Reply(200).
			JSON(map[string]string{"fname": "paypal", "url": "paypal-url"})
		gock.New("https://gock.com").
			Get("/payment/stripe/buttons/100").
			Reply(200).
			JSON(map[string]string{"fname": "paypal", "url": "paypal-url"})
		defer gock.Off()
		request, _ := http.NewRequest(http.MethodPost, "/payment/buttons", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		if writer.Code != 500 {
			t.Errorf("Status should be 500")
		}
	})

	t.Run("should returns X-Response-Time ", func(t *testing.T) {
		gock.New("https://gock.com").
			Get("/payment/paypal/buttons/100").
			Reply(200).
			JSON(map[string]string{"name": "paypal", "url": "paypal-url"})
		gock.New("https://gock.com").
			Get("/payment/stripe/buttons/100").
			Reply(200).
			JSON(map[string]string{"fname": "paypal", "url": "paypal-url"})
		defer gock.Off()
		request, _ := http.NewRequest(http.MethodPost, "/payment/buttons", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		header := writer.Header()["X-Response-Time"]

		if header == nil {
			t.Errorf("X-Response-Time missed")
		}
	})
}
