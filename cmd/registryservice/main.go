package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.Service{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		// 如果按任何键则继续往下执行，否则等待用户输入
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			return
		}
		// 关闭服务
		err = srv.Shutdown(ctx)
		if err != nil {
			return
		}
		cancel()
	}()

	<-ctx.Done()

	fmt.Println("Shutting down registry service.")
}
