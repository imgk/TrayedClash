package clash

import (
	"strconv"
	"github.com/Dreamacro/clash/hub/executor"
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

	instance := &Clash{
		Secret: conf.General.Secret,
		Server: conf.General.ExternalController,
	}
	return instance, nil
}

func (c *Clash) GetProxies() *sysproxy.ProxyConfig {
	general := executor.GetGeneral()
	return &sysproxy.ProxyConfig{
                Enable: true,
                Server: "127.0.0.1:" + strconv.Itoa(general.Port),
        }
}
