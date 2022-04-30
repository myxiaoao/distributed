package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, host, port string, registration registry.Registration,
	registerHandlersFunc func()) (context.Context, error) {
	// 注册服务
	registerHandlersFunc()
	// 启动服务
	ctx = startService(ctx, registration.ServiceName, host, port)
	err := registry.AddService(registration)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	// 重建一个具有取消功能 context
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		// 启动服务，如果有错误则打印错误
		log.Println(srv.ListenAndServe())
		// 取消服务注册
		err := registry.DelService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		// 取消
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		// 如果按任何键则继续往下执行，否则等待用户输入
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			return
		}
		err = registry.DelService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		// 关闭服务
		err = srv.Shutdown(ctx)
		if err != nil {
			return
		}
		cancel()
	}()

	return ctx
}
