package middleware

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(
	   prometheus.CounterOpts{
		  Namespace: "golang",
		  Name:      "request_counter_to_api",
		  Help:      "counts incoming requests to api",
	   })
  )

func Register() {
	prometheus.MustRegister(counter)
}

func MaskPromHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}