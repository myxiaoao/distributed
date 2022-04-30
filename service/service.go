package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string,
	registerHandlersFunc func()) (context.Context, error) {
	// 注册服务
	registerHandlersFunc()
	// 启动服务
	ctx = startService(ctx, serviceName, host, port)

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	// 重建一个具有取消功能 context
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		// 启动服务，如果有错误则打印错误
		log.Println(srv.ListenAndServe())
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
		// 关闭服务
		err = srv.Shutdown(ctx)
		if err != nil {
			return
		}
		cancel()
	}()

	return ctx
}
