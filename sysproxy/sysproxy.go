package sysproxy

var systemConfig = GetProxy()

type ProxyConfig struct {
	Enable bool
	Server string
}

func GetSystemConfig() *ProxyConfig {
	return systemConfig
}
