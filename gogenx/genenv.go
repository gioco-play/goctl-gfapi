package gogenx

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"strconv"
)

//go:embed env.tpl
var envTemplate string

func genEnv(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	service := api.Service
	host := "0.0.0.0"
	port := strconv.Itoa(defaultPort)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        ".env",
		templateName:    "envTemplate",
		category:        category,
		templateFile:    envTemplateFile,
		builtinTemplate: envTemplate,
		data: map[string]string{
			"serviceName": service.Name,
			"host":        host,
			"port":        port,
		},
	})
}
