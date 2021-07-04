package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	stdlog "log"
)

func main() {
	log.NewLog("./testlog")

	regInfo := registry.Registration{
		ServiceName: "Log",
		ServiceURL:  "localhost:8001",
	}

	ctx, err := service.Start(context.Background(), regInfo, log.RegisterHandler)
	if err != nil {
		stdlog.Println(err)
		return
	}

	<-ctx.Done()

	stdlog.Println("Log service stop.")
}
