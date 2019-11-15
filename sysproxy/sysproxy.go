package sysproxy

import (
	"strconv"
	"sync"

	"github.com/Dreamacro/clash/proxy"
)

var saveProxy = &sync.Once{}

func saveSystemProxy() {
	systemProxy := GetCurrentProxy()

	if systemProxy.Enable && systemProxy.Server == "127.0.0.1:"+strconv.Itoa(proxy.GetPorts().Port) {
		SystemProxy = &ProxyConfig{
			Enable: false,
			Server: ":80",
		}
	} else {
		SystemProxy = systemProxy
	}
}

// SystemProxy is ...
var SystemProxy = &ProxyConfig{}

// ProxyConfig is ...
type ProxyConfig struct {
	Enable bool
	Server string
}
