//go:generate go-bindata -pkg static -ignore .../.DS_Store -o files.go files/...

package static

import (
	"log"
	"net/http"
	"os"

	"github.com/elazarl/go-bindata-assetfs"
)

func ListenAndServe() {
	var handler http.Handler
	if info, err := os.Stat("static/files/"); err == nil && info.IsDir() {
		handler = http.FileServer(http.Dir("static/files/"))
	} else {
		handler = http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "files"})
	}

	log.Fatal(http.ListenAndServe("127.0.0.1:8780", handler))
}
