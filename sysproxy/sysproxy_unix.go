//+build linux darwin

package sysproxy

// GetCurrentProxy is ...
func GetCurrentProxy() *ProxyConfig {
	return &ProxyConfig{
		Enable: false,
		Server: ":80",
	}
}

// SetSystemProxy is ...
func SetSystemProxy(p *ProxyConfig) {
}
