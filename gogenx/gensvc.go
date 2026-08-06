package gogenx

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"

	dep "github.com/gioco-play/goctl-gfdep"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

const contextFilename = "service_context"

//go:embed context.tpl
var contextTemplate string

func genServiceContext(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, contextFilename)
	if err != nil {
		return err
	}

	middlewares := getMiddleware(api)

	imports := []string{"\"" + pathx.JoinPackages(rootPkg, configDir) + "\""}
	var fields, assigns []string
	for _, item := range middlewares {
		fields = append(fields, fmt.Sprintf("%s rest.Middleware", item))
		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		assigns = append(assigns, fmt.Sprintf("%s: middleware.New%s(redisClient, databasex).Handle,",
			item, strings.Title(name)))
	}
	if len(fields) > 0 {
		imports = append(imports,
			"\""+pathx.JoinPackages(rootPkg, middlewareDir)+"\"",
			fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL))
	}

	filePath := path.Join(dir, contextDir, filename+".go")
	if len(fields) > 0 && pathx.FileExists(filePath) {
		return patchExistingServiceContext(filePath, imports, fields, assigns)
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          contextDir,
		filename:        filename + ".go",
		templateName:    "contextTemplate",
		category:        category,
		templateFile:    contextTemplateFile,
		builtinTemplate: contextTemplate,
		data: map[string]string{
			"configImport":         strings.Join(imports, "\n\t"),
			"config":               "config.Config",
			"middleware":           strings.Join(fields, "\n"),
			"middlewareAssignment": strings.Join(assigns, "\n"),
		},
	})
}

// patchExistingServiceContext inserts the middleware wiring into an already
// generated service_context.go that genFile would otherwise skip untouched.
// On failure it downgrades to a warning listing what to paste in by hand,
// instead of failing the whole generation.
func patchExistingServiceContext(filePath string, imports, fields, assigns []string) error {
	applied, err := dep.PatchServiceContext(filePath, imports, fields, assigns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: %s: %v\n", filePath, err)
		fmt.Fprintf(os.Stderr, "請手動補上以下內容：\n")
		for _, l := range append(append(append([]string{}, imports...), fields...), assigns...) {
			fmt.Fprintf(os.Stderr, "  %s\n", l)
		}
		return nil
	}
	if len(applied) == 0 {
		return nil
	}

	fmt.Printf("%s 已更新：\n", filePath)
	for _, l := range applied {
		fmt.Printf("  + %s\n", l)
	}
	fmt.Println("提醒：internal/middleware 下對應的 Handle 目前是 passthrough，請填入邏輯")
	return nil
}
