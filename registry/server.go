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
		log.Printf("Add service [%s] at [%s]", r.ServiceName, r.ServiceURL)
		reg.add(r)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
