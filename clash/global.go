package clash 

import (
	"github.com/Dreamacro/clash/tunnel"
	A "github.com/Dreamacro/clash/adapters/outbound"
)

func (c *Clash) GetGlobalProxies() []string {
	p := []string{}

	for k, _ := range tunnel.Instance().Proxies() {
		switch k {
		case "GLOBAL", "REJECT", "DIRECT" :
		default:
			p = append(p, k)
		}
	}

	return p
}

func (c *Clash) GetGlobalNow() string {
	selector, _ := tunnel.Instance().Proxies()["GLOBAL"].(*A.Proxy).ProxyAdapter.(*A.Selector)

	return selector.Now()
}

func (c *Clash) SetGlobalNow(s string) error {
	selector, _ := tunnel.Instance().Proxies()["GLOBAL"].(*A.Proxy).ProxyAdapter.(*A.Selector)

	return selector.Set(s)
} 
