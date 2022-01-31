package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

func main() {
	tracer, err := newTracer()
	if err != nil {
		log.Fatal(err)
	}

	// We add the instrumented transport to the defaultClient
	// that comes with the zipkin-go library
	http.DefaultClient.Transport, err = zipkinhttp.NewTransport(
		tracer,
		zipkinhttp.TransportTrace(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandlerFactory(http.DefaultClient))
	r.HandleFunc("/foo", FooHandler)
	r.Use(zipkinhttp.NewServerMiddleware(
		tracer,
		zipkinhttp.SpanName("request")), // name for request span
	)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
