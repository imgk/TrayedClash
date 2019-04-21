package clash

import (
	"github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/config"
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
