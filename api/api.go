package api

import (
	"fmt"
	"log"
	"net/http"
)

type Params struct {
	Port string
}

func Create(params *Params) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Here's Johnny!")
	})

	NewPaymentHandler(mux)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", params.Port),
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}
