package gogenx

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

const (
	utilDir   = "utils"
	errorxDir = "utils/errorx"
)

//go:embed redislock.tpl
var redislockTemplate string

//go:embed tools.tpl
var toolsTemplate string
var errorxTemplate string

func genUtil(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	err := genRedislock(dir, cfg, api)
	if err != nil {
		return err
	}

	err = genTools(dir, cfg, api)
	if err != nil {
		return err
	}

	err = genErrorx(dir, cfg, api)
	if err != nil {
		return err
	}

	return nil
}

func genRedislock(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          utilDir,
		filename:        "redislock.go",
		templateName:    "redislockTemplate",
		category:        category,
		templateFile:    redislockTemplateFile,
		builtinTemplate: redislockTemplate,
		data: map[string]string{
			"serviceName": service.Name,
		},
	})
}

func genTools(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          utilDir,
		filename:        "tools.go",
		templateName:    "toolsTemplate",
		category:        category,
		templateFile:    toolsTemplateFile,
		builtinTemplate: toolsTemplate,
		data: map[string]string{
			"serviceName": service.Name,
		},
	})
}

func genErrorx(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          errorxDir,
		filename:        "errorx.go",
		templateName:    "errorxTemplate",
		category:        category,
		templateFile:    errorxTemplateFile,
		builtinTemplate: errorxTemplate,
		data: map[string]string{
			"serviceName": service.Name,
		},
	})
}
