package gogenx

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"strings"
)

const (
	respDir = "utils/respx"
)

//go:embed response.tpl
var respTemplate string

//go:embed respstate.tpl
var respstateTemplate string

func genResp(dir string, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {

	service := api.Service

	err := genRespState(dir, rootPkg, api)
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
			"serviceName":    service.Name,
			"ImportPackages": genRespImports(rootPkg),
		},
	})
}

func genRespState(dir string, rootPkg string, api *spec.ApiSpec) error {
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
			"serviceName":    service.Name,
			"ImportPackages": genRespImports(rootPkg),
		},
	})

}

func genRespImports(parentPkg string) string {
	imports := []string{
		fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, "utils/errorx")),
	}

	return strings.Join(imports, "\n\t")
}
