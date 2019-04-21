package clash

import (
	"strconv"

	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/imgk/TrayedClash/sysproxy"
)

func (c *Clash) GetProxies() *sysproxy.ProxyConfig {
	return &sysproxy.ProxyConfig{
        Enable: true,
        Server: "127.0.0.1:" + strconv.Itoa(proxy.GetPorts().Port),
    }
}

func (c *Clash) UpdateProxies(proxies map[string]constant.Proxy) error {
	for _, v := range tunnel.Instance().Proxies() {
		v.Destroy()
	}

	tunnel.Instance().UpdateProxies(proxies)

	return nil
}
