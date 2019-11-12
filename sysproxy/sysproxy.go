package sysproxy

import (
	"syscall"

	"golang.org/x/sys/windows/registry"
)

var SystemProxy = GetCurrentProxy()

type ProxyConfig struct {
	Enable bool
	Server string
}

func GetCurrentProxy() *ProxyConfig {
	k, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	defer k.Close()

	Enable, _, _ := k.GetIntegerValue("ProxyEnable")
	Server, _, _ := k.GetStringValue("ProxyServer")

	return &ProxyConfig{
		Enable: Enable > 0,
		Server: Server,
	}
}

func SetSystemProxy(p *ProxyConfig) {
	k, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	defer k.Close()

	if p.Enable {
		k.SetDWordValue("ProxyEnable", 0x00000001)
	} else {
		k.SetDWordValue("ProxyEnable", 0x00000000)
	}
	k.SetStringValue("ProxyServer", p.Server)

	func() {
		var wininet, _ = syscall.LoadLibrary("Wininet.dll")
		defer syscall.FreeLibrary(wininet)
		var internetSetOptionA, _ = syscall.GetProcAddress(wininet, "InternetSetOptionA")

		syscall.Syscall6(uintptr(internetSetOptionA), 4, uintptr(0), 0x0000005f, uintptr(0), 0x00000000, 0, 0)
		syscall.Syscall6(uintptr(internetSetOptionA), 4, uintptr(0), 0x00000025, uintptr(0), 0x00000000, 0, 0)
	}()
}
