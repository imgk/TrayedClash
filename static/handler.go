package static

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed gh-pages/140.4c90b0e4dc4d2542b082.js
//go:embed gh-pages/140.4c90b0e4dc4d2542b082.js.map
//go:embed gh-pages/272.43f7d3d4b29bad961dc5.js
//go:embed gh-pages/272.43f7d3d4b29bad961dc5.js.map
//go:embed gh-pages/354.07d4d292036733637e17.js
//go:embed gh-pages/354.07d4d292036733637e17.js.LICENSE.txt
//go:embed gh-pages/354.07d4d292036733637e17.js.map
//go:embed gh-pages/36.ba889d8c27f9ad2df62f.js
//go:embed gh-pages/36.ba889d8c27f9ad2df62f.js.map
//go:embed gh-pages/776.a0d9f3cb239ab5793543.js
//go:embed gh-pages/776.a0d9f3cb239ab5793543.js.map
//go:embed gh-pages/857.3187fe170ee00e40106d.js
//go:embed gh-pages/857.3187fe170ee00e40106d.js.map
//go:embed gh-pages/88.3c2f733e23b42b712d98.js
//go:embed gh-pages/88.3c2f733e23b42b712d98.js.map
//go:embed gh-pages/965.597108e0ccdfe9c47655.js
//go:embed gh-pages/965.597108e0ccdfe9c47655.js.LICENSE.txt
//go:embed gh-pages/965.597108e0ccdfe9c47655.js.map
//go:embed gh-pages/app.1141a20948b298f2a01c.css
//go:embed gh-pages/app.1141a20948b298f2a01c.css.map
//go:embed gh-pages/app.6706b8885424994ac6fe.js
//go:embed gh-pages/app.6706b8885424994ac6fe.js.map
//go:embed gh-pages/apple-touch-icon-precomposed.png
//go:embed gh-pages/chartjs.6115ab8a46716ce46acc.js
//go:embed gh-pages/chartjs.6115ab8a46716ce46acc.js.LICENSE.txt
//go:embed gh-pages/chartjs.6115ab8a46716ce46acc.js.map
//go:embed gh-pages/CNAME
//go:embed gh-pages/config.4667386daf3b0f4af026.css
//go:embed gh-pages/config.4667386daf3b0f4af026.css.map
//go:embed gh-pages/config.c78db5bb1c3ab25a375e.js
//go:embed gh-pages/config.c78db5bb1c3ab25a375e.js.map
//go:embed gh-pages/conns.1597b5ef9af1d66eaded.css
//go:embed gh-pages/conns.1597b5ef9af1d66eaded.css.map
//go:embed gh-pages/conns.f4d496fc2c5b3ae4c474.js
//go:embed gh-pages/conns.f4d496fc2c5b3ae4c474.js.map
//go:embed gh-pages/corejs.03c6fcb99c4847c87887.js
//go:embed gh-pages/corejs.03c6fcb99c4847c87887.js.map
//go:embed gh-pages/index.html
//go:embed gh-pages/libs.1b600ae622a5d20df4d5.js
//go:embed gh-pages/libs.1b600ae622a5d20df4d5.js.LICENSE.txt
//go:embed gh-pages/libs.1b600ae622a5d20df4d5.js.map
//go:embed gh-pages/logs.0784eb33531179dcd51f.js
//go:embed gh-pages/logs.0784eb33531179dcd51f.js.map
//go:embed gh-pages/logs.fdfa037875bf344e16eb.css
//go:embed gh-pages/logs.fdfa037875bf344e16eb.css.map
//go:embed gh-pages/open-sans-latin-400-normal.woff2
//go:embed gh-pages/open-sans-latin-700-normal.woff2
//go:embed gh-pages/proxies.2d1ba03fd1ea37eb898a.css
//go:embed gh-pages/proxies.2d1ba03fd1ea37eb898a.css.map
//go:embed gh-pages/proxies.5a3bcb921f085db99d99.js
//go:embed gh-pages/proxies.5a3bcb921f085db99d99.js.map
//go:embed gh-pages/report.html
//go:embed gh-pages/roboto-mono-latin-400-normal.woff2
//go:embed gh-pages/rules.4b07fbf893c0c51ef583.css
//go:embed gh-pages/rules.4b07fbf893c0c51ef583.css.map
//go:embed gh-pages/rules.a922219d35d4eb3c0f13.js
//go:embed gh-pages/rules.a922219d35d4eb3c0f13.js.map
//go:embed gh-pages/sw.js
//go:embed gh-pages/sw.js.map
//go:embed gh-pages/vendor.069567878b016e0c96b1.js
//go:embed gh-pages/vendor.069567878b016e0c96b1.js.LICENSE.txt
//go:embed gh-pages/vendor.069567878b016e0c96b1.js.map
//go:embed gh-pages/yacd-128.png
//go:embed gh-pages/yacd-64.png
//go:embed gh-pages/yacd.ico
//go:embed gh-pages/_headers
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
