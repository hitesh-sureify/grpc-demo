package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(
	   prometheus.CounterOpts{
		  Namespace: "golang",
		  Name:      "request_counter",
		  Help:      "counts incoming requests",
	   })
  )

func Register() {
	prometheus.MustRegister(counter)
}

func StartPrometheus(){
	go func() {
		pServer := http.NewServeMux()
		pServer.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", pServer)
	}()
}