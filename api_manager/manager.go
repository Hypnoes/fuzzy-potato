package api_manager

import (
	"net/http"

	apiservice "fuzzypotato/api_service"
)

type ServiceRegistry struct {
	Service map[string]apiservice.ServiceConfig
}

func (r *ServiceRegistry) Register(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.FormValue("name")
	queue := req.FormValue("queue")
	methods := req.FormValue("methods")

	r.Service[name] = apiservice.ServiceConfig{
		Name:   name,
		Queue:  queue,
		Method: map[string]string{"methods": methods},
	}

	w.WriteHeader(200)
	w.Write([]byte("Registered"))
}

func (r *ServiceRegistry) Lookup(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.FormValue("name")

	if service, ok := r.Service[name]; ok {
		w.WriteHeader(200)
		w.Write([]byte(service.Queue))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Service not found"))
	}
}
