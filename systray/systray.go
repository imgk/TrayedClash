package systray

import (
	"runtime"
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"

	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"

	"github.com/imgk/TrayedClash/icon"
	"github.com/imgk/TrayedClash/sysproxy"
)

func init() {
	runtime.LockOSThread()
}

// Run is ...
func Run() {
	systray.RunWithAppWindow("Clash", 960, 640, onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Clash")
	systray.SetTooltip("A Rule-based Tunnel in Go")

	mTitle := systray.AddMenuItem("Clash - A Rule-Based Tunnel", "")
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

		SavedPort := proxy.GetPorts().Port
		for {
			<-t.C

			switch tunnel.Mode() {
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
				p := proxy.GetPorts().Port
				if SavedPort != p {
					SavedPort = p
					err := sysproxy.SetSystemProxy(
						&sysproxy.ProxyConfig{
							Enable: true,
							Server: "127.0.0.1:" + strconv.Itoa(SavedPort),
						})
					if err != nil {
						continue
					}
				}
			}

			p, err := sysproxy.GetCurrentProxy()
			if err != nil {
				continue
			}

			if p.Enable && p.Server == "127.0.0.1:"+strconv.Itoa(SavedPort) {
				if mEnabled.Checked() {
				} else {
					mEnabled.Check()
				}
			} else {
				if mEnabled.Checked() {
					mEnabled.Uncheck()
				} else {
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-mTitle.ClickedCh:
			case <-mGlobal.ClickedCh:
				tunnel.SetMode(tunnel.Global)
			case <-mRule.ClickedCh:
				tunnel.SetMode(tunnel.Rule)
			case <-mDirect.ClickedCh:
				tunnel.SetMode(tunnel.Direct)
			case <-mEnabled.ClickedCh:
				if mEnabled.Checked() {
					err := sysproxy.SetSystemProxy(sysproxy.GetSavedProxy())
					if err != nil {
					} else {
						mEnabled.Uncheck()
					}
				} else {
					err := sysproxy.SetSystemProxy(
						&sysproxy.ProxyConfig{
							Enable: true,
							Server: "127.0.0.1:" + strconv.Itoa(proxy.GetPorts().Port),
						})
					if err != nil {
					} else {
						mEnabled.Check()
					}
				}
			case <-mURL.ClickedCh:
				switch runtime.GOOS {
				case "darwin":
					systray.ShowAppWindow("http://clash.razord.top/")
				case "windows":
					err := open.Run("http://yacd.haishan.me/")
					if err != nil {
					}
				default:
					systray.ShowAppWindow("http://clash.razord.top/")
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	for {
		err := sysproxy.SetSystemProxy(sysproxy.GetSavedProxy())
		if err != nil {
			continue
		} else {
			break
		}
	}
}
