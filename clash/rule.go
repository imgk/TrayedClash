package clash 

import (
	"github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/tunnel"
)

func (c *Clash) UpdateRules(rules []constant.Rule) error {
	tunnel.Instance().UpdateRules(rules)

	return nil
}
