package middleware

import (
	"net/http"
	//"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


	var Incoming_api_req_counter = prometheus.NewCounter(
	   prometheus.CounterOpts{
		  Namespace: "golang",
		  Name:      "request_counter_to_api",
		  Help:      "counts incoming requests to api",
	   })
	
	var Emp_get_fail_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "get_employee_fail_ctr",
		Help:      "counts failure to fetch employee",
	})

	var Emp_create_fail_counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "create_employee_fail_ctr",
			Help:      "counts failure to create employee",
		})
	
	var Emp_update_fail_counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "update_employee_fail_ctr",
			Help:      "counts failure to update employee",
		})
	
	var Emp_delete_fail_counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "delete_employee_fail_ctr",
			Help:      "counts failure to delete employee",
		})

func Register() {
	prometheus.MustRegister(Incoming_api_req_counter)
	prometheus.MustRegister(Emp_get_fail_counter)
	prometheus.MustRegister(Emp_create_fail_counter)
	prometheus.MustRegister(Emp_update_fail_counter)
	prometheus.MustRegister(Emp_delete_fail_counter)
}

func MaskPromHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}