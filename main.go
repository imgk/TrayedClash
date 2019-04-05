package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub"

	log "github.com/sirupsen/logrus"

	"github.com/imgk/TrayedClash/systray"
)

var (
	version bool
	homedir string
)

func init() {
	flag.StringVar(&homedir, "d", "", "set configuration directory")
	flag.BoolVar(&version, "v", false, "show current version of clash")
	flag.Parse()
}

func main() {
	if version {
		fmt.Printf("Clash %s %s %s %s\n", C.Version, runtime.GOOS, runtime.GOARCH, C.BuildTime)
		return
	}

	// enable tls 1.3 and remove when go 1.13
	os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",tls13=1")

	if homedir != "" {
		if !filepath.IsAbs(homedir) {
			currentDir, _ := os.Getwd()
			homedir = filepath.Join(currentDir, homedir)
		}
		C.SetHomeDir(homedir)
	} else {
		if runtime.GOOS == "windows" {
			currentDir, _ := os.Getwd()
			C.SetHomeDir(currentDir)
		}
	}

	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalf("Initial configuration directory error: %s", err.Error())
	}

	if err := hub.Parse(); err != nil {
		log.Fatalf("Parse config error: %s", err.Error())
	}

	systray.Run()
}
