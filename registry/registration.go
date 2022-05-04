package registry

type Registration struct {
	// 服务名称
	ServiceName ServiceName
	// 服务地址
	ServiceURL string
	// 当前服务依赖其他服务
	RequiredServices []ServiceName
	// 服务更新地址
	ServiceUpdateURL string
	// 服务心跳检查地址
	HeartbeatURL string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portal")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
