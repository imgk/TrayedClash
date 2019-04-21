package systray

import (
	"runtime"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"

	"github.com/imgk/TrayedClash/clash"
	"github.com/imgk/TrayedClash/icon"
	"github.com/imgk/TrayedClash/sysproxy"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	switch runtime.GOOS {
	case "darwin":
		systray.SetTitle("")
	default:
		systray.SetTitle("Clash - A rule-based tunnel in Go")
	}
	systray.SetTooltip("Clash")

	mGlobal := systray.AddMenuItem("Global", "Set as Global")
	mRule := systray.AddMenuItem("Rule", "Set as Rule")
	mDirect := systray.AddMenuItem("Direct", "Set as Direct")

	systray.AddSeparator()

	mGlobalProxies := make(map[string]*systray.MenuItem)
	for _, v := range clash.GetInstance().GetGlobalProxies() {
		mGlobalProxies[v] = systray.AddMenuItem(v, "Set Global as "+v)
	}

	systray.AddSeparator()

	mEnabled := systray.AddMenuItem("Set as System Proxy", "Turn on/off Proxy")
	mUpConfig := systray.AddMenuItem("Update Remote Config", "Update Remote Config")
	mUrl := systray.AddMenuItem("Open Clash Dashboard", "Open Clash Dashboard")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Exit", "Quit Clash")

	changeMode := func(mode clash.Mode) {
		switch mode {
		case clash.Global:
			mGlobal.Check()
			mRule.Uncheck()
			mDirect.Uncheck()
		case clash.Rule:
			mGlobal.Uncheck()
			mRule.Check()
			mDirect.Uncheck()
		case clash.Direct:
			mGlobal.Uncheck()
			mRule.Uncheck()
			mDirect.Check()
		}
	}

	changeGlobal := func(mode clash.Mode) {
		if mode == clash.Global {
			for k, v := range mGlobalProxies {
				v.Enable()
				if k == clash.GetInstance().GetGlobalNow() {
					v.Check()
				} else {
					v.Uncheck()
				}
			}
		} else {
			for _, v := range mGlobalProxies {
				v.Disable()
			}
		}
	}

	changeMode(clash.GetInstance().GetMode())
	changeGlobal(clash.GetInstance().GetMode())

	for name, item := range mGlobalProxies {
		var ch = make(chan int)
		go func(proxy string, menu *systray.MenuItem) {
			ch <- 1
			for {
				<-menu.ClickedCh
				if clash.GetInstance().SetGlobalNow(proxy) != nil {
					continue
				}
				changeGlobal(clash.GetInstance().GetMode())
			}
		}(name, item)
		// wait for goroutine to create
		<-ch
		close(ch)
	}

	go func() {
		for {
			select {
			case <-mGlobal.ClickedCh:
				changeMode(clash.Global)
				clash.GetInstance().SetMode(clash.Global)
				changeGlobal(clash.GetInstance().GetMode())
			case <-mRule.ClickedCh:
				changeMode(clash.Rule)
				clash.GetInstance().SetMode(clash.Rule)
				changeGlobal(clash.GetInstance().GetMode())
			case <-mDirect.ClickedCh:
				changeMode(clash.Direct)
				clash.GetInstance().SetMode(clash.Direct)
				changeGlobal(clash.GetInstance().GetMode())
			case <-mEnabled.ClickedCh:
				if mEnabled.Checked() {
					mEnabled.Uncheck()
					sysproxy.SetProxy(sysproxy.GetSystemConfig())
				} else {
					mEnabled.Check()
					sysproxy.SetProxy(clash.GetInstance().GetProxies())
				}
			case <-mUpConfig.ClickedCh:
				clash.GetInstance().UpdateConfigFromRemote()
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
	sysproxy.SetProxy(sysproxy.GetSystemConfig())
}
