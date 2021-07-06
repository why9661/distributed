package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegistryHandler{})

	var srv http.Server
	srv.Addr = registry.RegistryHost + ":" + registry.RegistryPort

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		log.Println("Registry Service start.Please press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
	}()

	<-ctx.Done()
	log.Println("Registry Service Stop.")
}
