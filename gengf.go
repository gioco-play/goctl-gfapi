package main

import (
	"github.com/gioco-play/goctl-gfapi/gogenx"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"strings"
)

func main() {
	plugin_, err := plugin.NewPlugin()
	if err != nil {
		panic(err)
	}

	if len(plugin_.Style) == 0 {
		plugin_.Style = "gozero"
	}

	err = gogenx.DoGenProject(plugin_.ApiFilePath, plugin_.Dir, strings.TrimSpace(plugin_.Style))
	if err != nil {
		panic(err)
	}
}
