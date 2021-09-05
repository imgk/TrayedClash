package sysproxy

import (
	"strconv"

	proxy "github.com/Dreamacro/clash/listener"
)

// ProxyConfig is ...
type ProxyConfig struct {
	Enable bool
	Server string
}

// SavedProxy is ...
var SavedProxy *ProxyConfig

func (c *ProxyConfig) String() string {
	if c == nil {
		return "nil"
	}

	if c.Enable {
		return "Enabled: True" + "; Server: " + c.Server
	}

	return "Enabled: False" + "; Server: " + c.Server
}

// GetSavedProxy is ...
func GetSavedProxy() *ProxyConfig {
	if SavedProxy == nil {
		err := func() error {
			p, err := GetCurrentProxy()
			if err != nil {
				return err
			}

			if p.Enable && p.Server == "127.0.0.1:"+strconv.Itoa(proxy.GetPorts().Port) {
				SavedProxy = &ProxyConfig{
					Enable: false,
					Server: ":80",
				}
			} else {
				SavedProxy = p
			}

			return nil
		}()
		if err != nil {
			SavedProxy = &ProxyConfig{
				Enable: false,
				Server: ":80",
			}
			return SavedProxy
		}

		return SavedProxy
	}

	return SavedProxy
}
