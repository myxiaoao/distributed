package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stLog "log"
)

// 日志服务
func main() {
	// 初始化自定义日志
	log.Run("./distributed.log")
	// 初始化主机，地址
	host, port := "localhost", "8002"
	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddr,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddr + "/services",
		HeartbeatURL:     serviceAddr + "/heartbeat",
	}
	// 启动服务
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		// 如果有错误，则打印日志。此时自定日志服务还未启动，使用系统的日志打印。
		stLog.Fatalln(err)
	}
	// 等待信号
	// 如果启动出现错误
	// 或者手动停止
	<-ctx.Done()

	fmt.Println("Shutting down log service.")
}
