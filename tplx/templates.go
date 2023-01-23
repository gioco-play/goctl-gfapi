package tplx

import (
	"fmt"
	"github.com/gioco-play/goctl-gfapi/gogenx"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/tools/goctl/api/apigen"
	apinew "github.com/zeromicro/go-zero/tools/goctl/api/new"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const templateParentPath = "/"

// genTemplates writes the latest template text into file which is not exists
func genTemplatesX(_ *cobra.Command, _ []string) error {
	path := varStringHome
	if len(path) != 0 {
		pathx.RegisterGoctlHome(path)
	}

	if err := errorx.Chain(
		func() error {
			return gogenx.GenTemplates()
		},
		func() error {
			return apigen.GenTemplates()
		},
		func() error {
			return apinew.GenTemplates()
		},
	); err != nil {
		return err
	}

	dir, err := pathx.GetTemplateDir(templateParentPath)
	if err != nil {
		return err
	}

	abs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	fmt.Printf("Templates are generated in %s, %s\n", aurora.Green(abs),
		aurora.Red("edit on your risk!"))

	return nil
}

// cleanTemplates deletes all templates
func cleanTemplatesX(_ *cobra.Command, _ []string) error {
	path := varStringHome
	if len(path) != 0 {
		pathx.RegisterGoctlHome(path)
	}

	err := errorx.Chain(
		func() error {
			return gogenx.Clean()
		},
		func() error {
			return apigen.Clean()
		},
		func() error {
			return apinew.Clean()
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", aurora.Green("templates are cleaned!"))
	return nil
}

// updateTemplates writes the latest template text into file,
// it will delete the older templates if there are exists
func updateTemplatesX(_ *cobra.Command, _ []string) (err error) {
	path := varStringHome
	category := varStringCategory
	if len(path) != 0 {
		pathx.RegisterGoctlHome(path)
	}

	defer func() {
		if err == nil {
			fmt.Println(aurora.Green(fmt.Sprintf("%s template are update!", category)).String())
		}
	}()
	switch category {
	case gogenx.Category():
		return gogenx.Update()
	case apigen.Category():
		return apigen.Update()
	case apinew.Category():
		return apinew.Update()
	default:
		err = fmt.Errorf("unexpected category: %s", category)
		return
	}
}

// revertTemplates will overwrite the old template content with the new template
func revertTemplatesX(_ *cobra.Command, _ []string) (err error) {
	path := varStringHome
	category := varStringCategory
	filename := varStringName
	if len(path) != 0 {
		pathx.RegisterGoctlHome(path)
	}

	defer func() {
		if err == nil {
			fmt.Println(aurora.Green(fmt.Sprintf("%s template are reverted!", filename)).String())
		}
	}()
	switch category {
	case gogenx.Category():
		return gogenx.RevertTemplate(filename)
	case apigen.Category():
		return apigen.RevertTemplate(filename)
	case apinew.Category():
		return apinew.RevertTemplate(filename)
	default:
		err = fmt.Errorf("unexpected category: %s", category)
		return
	}
}
