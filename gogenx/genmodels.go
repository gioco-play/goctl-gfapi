package gogenx

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

const (
	modelsDir = "models"
)

//go:embed models.tpl
var modelsTemplate string

func genModels(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          modelsDir,
		filename:        "models.go",
		templateName:    "modelsTemplate",
		category:        category,
		templateFile:    modelsTemplateFile,
		builtinTemplate: modelsTemplate,
		data: map[string]string{
			"serviceName": service.Name,
		},
	})
}
