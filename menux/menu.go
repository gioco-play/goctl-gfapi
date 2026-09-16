// Package menux provides the interactive "menu" subcommand, which in-process
// wraps the two API actions that used to live in the standalone
// goctl-gfmenu CLI (generate an API service, add a dependency to one).
//
// ADR 結論（不要在後續變動時忘記）：
//
//  1. 生成功能不能省略、不能交給 makefile：makefile 的 api target 本身是範本
//     產物，首次生成時它還不存在。此選單是首次生成唯一的入口。
//  2. 只設 gogenx.VarStringHome，不要改設 gogenx.VarStringRemote：
//     gogenx/gen.go 的 GoCommand 在 remote clone 失敗時會吞掉錯誤並靜默改用
//     內建範本，產出一個形似但並非本組織的骨架；而 gogenx/adddep.go 的
//     AddDep 走 remote 路徑時失敗則會直接 return err，兩條路徑並不對稱。
//     這裡改為自行 CloneIntoGitHome、檢查錯誤、再透過 --home 傳入範本路徑，
//     就是為了繞開這個不對稱行為。
package menux

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/gioco-play/goctl-gfapi/gogenx"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/util"
)

const templateRepo = "https://github.com/gioco-play/gf-template"

const (
	labelGen    = "產生 API 服務"
	labelAddDep = "API 加依賴"
)

// Cmd is the top-level "menu" subcommand. It intentionally takes no flags:
// every value it needs is discovered by scanning the current directory or
// asked interactively.
var Cmd = &cobra.Command{
	Use: "menu",
	// 走錯目錄之類的預期錯誤不需要 usage；main() 已負責印訊息。
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Interactive menu to generate an API service or add a dependency",
	RunE:          run,
}

func run(_ *cobra.Command, _ []string) error {
	apiFiles, _ := filepath.Glob("*.api")
	sort.Strings(apiFiles)
	if len(apiFiles) == 0 {
		return errors.New("找不到 *.api，請在服務的 api/ 目錄下執行")
	}

	genLabel := labelGen
	if len(apiFiles) == 1 {
		genLabel = fmt.Sprintf("%s    %s", labelGen, apiFiles[0])
	}

	var choice string
	if err := askOrCancel(survey.AskOne(&survey.Select{
		Message: "請選擇動作：",
		Options: []string{genLabel, labelAddDep},
	}, &choice)); err != nil || choice == "" {
		return err
	}

	isGen := choice == genLabel

	var apiFile string
	if isGen {
		apiFile = apiFiles[0]
		if len(apiFiles) > 1 {
			if err := askOrCancel(survey.AskOne(&survey.Select{
				Message: "偵測到多個候選檔，請選擇：",
				Options: apiFiles,
			}, &apiFile)); err != nil || apiFile == "" {
				return err
			}
		}
	}

	tplPath, err := util.CloneIntoGitHome(templateRepo, "")
	if err != nil {
		return err
	}

	var deps []string
	if !isGen {
		matches, _ := filepath.Glob(filepath.Join(tplPath, "deps", "*.tpl"))
		names := make([]string, len(matches))
		for i, m := range matches {
			names[i] = strings.TrimSuffix(filepath.Base(m), ".tpl")
		}
		sort.Strings(names)

		if err := askOrCancel(survey.AskOne(&survey.MultiSelect{
			Message: "請選擇要加入的依賴：",
			Options: names,
		}, &deps)); err != nil {
			return err
		}
		if len(deps) == 0 {
			return nil
		}
	}

	// 印出的是解析後的真實值（實際 --home 快取路徑、實際勾選的依賴），不是
	// 漂亮化的 --remote 形式：這條命令是唯一的肉眼防線，全域變數設錯名字或
	// 漏設不會編譯失敗只會靜默跑錯。
	var equivalent string
	if isGen {
		equivalent = fmt.Sprintf("goctl-gfapi go --api %s --dir . --home %s --style gozero", apiFile, tplPath)
	} else {
		equivalent = fmt.Sprintf("goctl-gfapi add-dep --name %s --home %s", strings.Join(deps, ","), tplPath)
	}
	fmt.Printf("\n即將執行：\n\n  %s\n\n", equivalent)

	confirm := false
	if err := askOrCancel(survey.AskOne(&survey.Confirm{Message: "執行？", Default: false}, &confirm)); err != nil {
		return err
	}
	if !confirm {
		return nil
	}

	if isGen {
		gogenx.VarStringAPI = apiFile
		gogenx.VarStringDir = "."
		gogenx.VarStringHome = tplPath
		gogenx.VarStringStyle = "gozero"
		return gogenx.GoCommand(nil, nil)
	}

	gogenx.VarStringName = strings.Join(deps, ",")
	gogenx.VarStringHome = tplPath
	return gogenx.AddDep(nil, nil)
}

// askOrCancel treats Ctrl-C during a survey prompt as a user-initiated
// cancellation (not a failure), returning nil so gengf.go's main() exits
// cleanly. Any other prompt error is returned as-is.
func askOrCancel(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, terminal.InterruptErr) {
		return nil
	}
	return err
}
