//+build linux darwin

package sysproxy

func GetProxy() *ProxyConfig {
	return &ProxyConfig {
		Enable: false,
		Server: "127.0.0.1:7890",
	}
}

func SetProxy(p *ProxyConfig) {
}
