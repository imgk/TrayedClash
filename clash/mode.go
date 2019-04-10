package clash

import (
	"github.com/Dreamacro/clash/tunnel"
)

type Mode tunnel.Mode

const (
	Global = Mode(tunnel.Global)
	Rule = Mode(tunnel.Rule)
	Direct = Mode(tunnel.Direct)
)

func (c *Clash) GetMode() Mode {
	return Mode(tunnel.Instance().Mode())
}

func (c *Clash) SetMode(mode Mode) {
	tunnel.Instance().SetMode(tunnel.Mode(mode))
}
