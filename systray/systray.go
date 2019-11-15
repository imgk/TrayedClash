package systray

import (
	"strconv"
	"time"

	"github.com/getlantern/systray"

	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"

	"github.com/imgk/TrayedClash/icon"
	"github.com/imgk/TrayedClash/sysproxy"
)

// Run is ...
func Run() {
	systray.RunWithAppWindow("Clash", 900, 600, onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Clash")
	systray.SetTooltip("A Rule-based Tunnel in Go")

	mTemp := systray.AddMenuItem("", "")
	mTest := systray.AddMenuItem("Clash - A Rule-based Tunnel", "")
	systray.AddSeparator()

	mGlobal := systray.AddMenuItem("Global", "Set as Global")
	mRule := systray.AddMenuItem("Rule", "Set as Rule")
	mDirect := systray.AddMenuItem("Direct", "Set as Direct")

	systray.AddSeparator()

	mEnabled := systray.AddMenuItem("Set as System Proxy", "Turn on/off Proxy")
	mURL := systray.AddMenuItem("Open Clash Dashboard", "Open Clash Dashboard")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Exit", "Quit Clash")

	go func() {
		t := time.NewTicker(time.Duration(time.Second))
		defer t.Stop()

		savedPort := proxy.GetPorts().Port
		for {
			<-t.C

			switch tunnel.Instance().Mode() {
			case tunnel.Global:
				if mGlobal.Checked() {
				} else {
					mGlobal.Check()
					mRule.Uncheck()
					mDirect.Uncheck()
				}
			case tunnel.Rule:
				if mRule.Checked() {
				} else {
					mGlobal.Uncheck()
					mRule.Check()
					mDirect.Uncheck()
				}
			case tunnel.Direct:
				if mDirect.Checked() {
				} else {
					mGlobal.Uncheck()
					mRule.Uncheck()
					mDirect.Check()
				}
			}

			if mEnabled.Checked() {
				clashPort := proxy.GetPorts().Port
				if savedPort != clashPort {
					savedPort = clashPort
					sysproxy.SetSystemProxy(
						&sysproxy.ProxyConfig{
							Enable: true,
							Server: "127.0.0.1:" + strconv.Itoa(savedPort),
						})
				}
			}

			systemProxy := sysproxy.GetCurrentProxy()
			if systemProxy.Enable && systemProxy.Server == "127.0.0.1:"+strconv.Itoa(savedPort) {
				if mEnabled.Checked() {
				} else {
					mEnabled.Check()
				}
			} else {
				if mEnabled.Checked() {
					mEnabled.Uncheck()
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-mTemp.ClickedCh:
			case <-mTest.ClickedCh:
			case <-mGlobal.ClickedCh:
				tunnel.Instance().SetMode(tunnel.Global)
			case <-mRule.ClickedCh:
				tunnel.Instance().SetMode(tunnel.Rule)
			case <-mDirect.ClickedCh:
				tunnel.Instance().SetMode(tunnel.Direct)
			case <-mEnabled.ClickedCh:
				if mEnabled.Checked() {
					mEnabled.Uncheck()
					sysproxy.SetSystemProxy(sysproxy.SystemProxy)
				} else {
					mEnabled.Check()
					sysproxy.SetSystemProxy(
						&sysproxy.ProxyConfig{
							Enable: true,
							Server: "127.0.0.1:" + strconv.Itoa(proxy.GetPorts().Port),
						})
				}
			case <-mURL.ClickedCh:
				systray.ShowAppWindow("http://clash.razord.top/")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	sysproxy.SetSystemProxy(sysproxy.SystemProxy)
}
