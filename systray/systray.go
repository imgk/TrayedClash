package systray

import(
	"runtime"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"github.com/imgk/TrayedClash/icon"
	"github.com/imgk/TrayedClash/sysproxy"
	"github.com/imgk/TrayedClash/clash"
)

var systemConfig = sysproxy.GetProxy()
var instance, _ = clash.NewClash()

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	if runtime.GOOS == "darwin" {
		systray.SetTitle("")
	} else if runtime.GOOS == "windows" {
		systray.SetTitle("Clash - A rule-based tunnel in Go")
	} else {
		systray.SetTitle("Clash - A rule-based tunnel in Go")
	}
	systray.SetTooltip("Clash")

	go func() {
		mEnabled := systray.AddMenuItem("Set as System Proxy", "Turn on/off Proxy")
		mUrl := systray.AddMenuItem("Clash Dashboard", "Open Clash Dashboard")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Exit", "Quit Clash")

		for {
			select {
			case <-mEnabled.ClickedCh:
				if mEnabled.Checked() {
					mEnabled.Uncheck()
					sysproxy.SetProxy(systemConfig)
				} else {
					mEnabled.Check()
					sysproxy.SetProxy(instance.GetProxies())
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
}
