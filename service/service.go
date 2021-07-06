package service

import (
	"context"
	"distributed/registry"
	"fmt"
	stdlog "log"
	"net/http"
)

func Start(ctx context.Context, r registry.Registration, registerHandler func()) (context.Context, error) {
	registerHandler()

	ctx, err := startService(ctx, r)
	if err != nil {
		return nil, err
	}

	return ctx, err
}

func startService(ctx context.Context, r registry.Registration) (context.Context, error) {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server

	srv.Addr = r.ServiceURL

	go func() {
		registry.RegisterService(r)
		stdlog.Println(srv.ListenAndServe())
		registry.RemoveService(r)
		cancel()
	}()

	go func() {
		stdlog.Printf("%s service start.Please press any key to stop.", r.ServiceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
	}()

	return ctx, nil
}
