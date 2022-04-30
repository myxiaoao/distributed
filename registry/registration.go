package registry

type Registration struct {
	// 服务名称
	ServiceName ServiceName
	// 服务地址
	ServiceURL string
}

type ServiceName string

const (
	LogService = ServiceName("LogService")
)
