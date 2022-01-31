package main

import (
	"bytes"
	"net/http"
)

func HomeHandlerFactory(client *http.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := bytes.NewBufferString("")
		res, err := client.Post("http://example.com", "application/json", body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if res.StatusCode > 399 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
