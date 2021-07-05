package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type registry struct {
	registrations []Registration
	m             *sync.Mutex
}

func (r *registry) add(reg Registration) {
	r.m.Lock()
	r.registrations = append(r.registrations, reg)
	r.m.Unlock()
}

func (r *registry) remove(reg Registration) {
	r.m.Lock()
	for i := range r.registrations {
		if r.registrations[i] == reg {
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			r.m.Unlock()
			return
		}
	}
	r.m.Unlock()
}

var reg = registry{make([]Registration, 0), new(sync.Mutex)}

type RegistryHandler struct{}

func (h *RegistryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received.")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registration
		if err := dec.Decode(&r); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		reg.add(r)
		log.Printf("Add service [%s] at [%s]", r.ServiceName, r.ServiceURL)
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		dec := json.NewDecoder(r.Body)
		var r Registration
		if err := dec.Decode(&r); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		reg.remove(r)
		log.Printf("Stop service [%s] at [%s]", r.ServiceName, r.ServiceURL)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
