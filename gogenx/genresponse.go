package gogenx

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

const (
	respDir = "internal/respx"
)

//go:embed response.tpl
var respTemplate string

//go:embed respstate.tpl
var respstateTemplate string

func genResp(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	service := api.Service

	err := genRespState(dir, cfg, api)
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          respDir,
		filename:        "respx.go",
		templateName:    "respTemplate",
		category:        category,
		templateFile:    respTemplateFile,
		builtinTemplate: respTemplate,
		data: map[string]string{
			"serviceName": service.Name,
		},
	})
}

func genRespState(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          respDir,
		filename:        "respstate.go",
		templateName:    "respstateTemplate",
		category:        category,
		templateFile:    respstateTemplateFile,
		builtinTemplate: respstateTemplate,
		data: map[string]string{
			"serviceName": service.Name,
		},
	})

}
