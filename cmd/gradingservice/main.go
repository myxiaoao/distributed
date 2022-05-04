package main

import (
	"context"
	"distributed/grades"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stLog "log"
)

func main() {
	host, port := "localhost", "8003"
	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddr,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddr + "/services",
		HeartbeatURL:     serviceAddr + "/heartbeat",
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers,
	)
	if err != nil {
		stLog.Fatalln(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()

	fmt.Println("Shutting down grading service.")
}
