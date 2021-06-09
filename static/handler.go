package static

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)
//go:embed gh-pages/CNAME
//go:embed gh-pages/_headers
//go:embed gh-pages/apple-touch-icon-precomposed.png
//go:embed gh-pages/index.html
//go:embed gh-pages/manifest.webmanifest
//go:embed gh-pages/registerSW.js
//go:embed gh-pages/sw.js
//go:embed gh-pages/yacd-128.png
//go:embed gh-pages/yacd-64.png
//go:embed gh-pages/yacd.ico
//go:embed gh-pages/assets/*
var content embed.FS

func init() {
	addr := "127.0.0.1:8780"
	if e := os.Getenv("ADMIN_ADDR"); e != "" {
		addr = e
	}
	fsys, err := fs.Sub(content, "gh-pages")
	if err != nil {
		log.Panic(err)
	}
	go func(fsys fs.FS, addr string) {
		if err := http.ListenAndServe(addr, http.FileServer(http.FS(fsys))); err != nil {
			log.Panicln(err)
		}
	}(fsys, addr)
}
