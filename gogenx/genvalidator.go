package gogenx

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

const (
	validatorDir = "internal/util"
)

//go:embed validator.tpl
var validatorTemplate string

func genValidator(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          validatorDir,
		filename:        "validator.go",
		templateName:    "validatorTemplate",
		category:        category,
		templateFile:    validatorTemplateFile,
		builtinTemplate: validatorTemplate,
		data: map[string]string{
			"serviceName": service.Name,
		},
	})
}
