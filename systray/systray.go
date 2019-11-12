package systray

import (
	"strconv"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"

	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"

	"github.com/imgk/TrayedClash/icon"
	"github.com/imgk/TrayedClash/sysproxy"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	go func() {
		systray.SetIcon(icon.Data)
		systray.SetTitle("Clash")
		systray.SetTooltip("A rule-based tunnel in Go")

		mTemp := systray.AddMenuItem("", "")
		mTest := systray.AddMenuItem("Clash - A Rule-based Tunnel", "")
		systray.AddSeparator()

		mGlobal := systray.AddMenuItem("Global", "Set as Global")
		mRule := systray.AddMenuItem("Rule", "Set as Rule")
		mDirect := systray.AddMenuItem("Direct", "Set as Direct")

		systray.AddSeparator()

		mEnabled := systray.AddMenuItem("Set as System Proxy", "Turn on/off Proxy")
		mUrl := systray.AddMenuItem("Open Clash Dashboard", "Open Clash Dashboard")

		systray.AddSeparator()

		mQuit := systray.AddMenuItem("Exit", "Quit Clash")

		switch tunnel.Instance().Mode() {
		case tunnel.Global:
			mGlobal.Check()
		case tunnel.Rule:
			mRule.Check()
		case tunnel.Direct:
			mDirect.Check()
		}

		for {
			select {
			case <-mTemp.ClickedCh:
			case <-mTest.ClickedCh:
			case <-mGlobal.ClickedCh:
				mGlobal.Check()
				mRule.Uncheck()
				mDirect.Uncheck()
				tunnel.Instance().SetMode(tunnel.Global)
			case <-mRule.ClickedCh:
				mGlobal.Uncheck()
				mRule.Check()
				mDirect.Uncheck()
				tunnel.Instance().SetMode(tunnel.Rule)
			case <-mDirect.ClickedCh:
				mGlobal.Uncheck()
				mRule.Uncheck()
				mDirect.Check()
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
			case <-mUrl.ClickedCh:
				open.Run("http://clash.razord.top")
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
