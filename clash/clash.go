package clash

import (
	"strconv"

	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/config"

	"github.com/imgk/TrayedClash/sysproxy"
)

type Clash struct {
	Secret string
	Server string
}

func NewClash() (*Clash, error) {
	conf, err := config.Parse(constant.Path.Config())
	if err != nil {
		return nil, err
	}

	return &Clash{
		Secret: conf.General.Secret,
		Server: conf.General.ExternalController,
	}, nil
}

func (c *Clash) GetProxies() *sysproxy.ProxyConfig {
	return &sysproxy.ProxyConfig{
        Enable: true,
        Server: "127.0.0.1:" + strconv.Itoa(proxy.GetPorts().Port),
    }
}
