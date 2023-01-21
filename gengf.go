package main

import (
	"fmt"
	"github.com/gioco-play/goctl-gfapi/gogenx"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const (
	codeFailure = 1
	dash        = "-"
	doubleDash  = "--"
	assign      = "="
)

var (
	rootCmd = &cobra.Command{
		Use:   "goctl-gfapi",
		Short: "A cli tool to generate go-zero code",
		Long: "A cli tool to generate api, zrpc, model code\n\n" +
			"GitHub: https://github.com/zeromicro/go-zero\n" +
			"Site:   https://go-zero.dev",
	}

	goCmd = &cobra.Command{
		Use:   "go",
		Short: "Generate go files for provided api in api file",
		RunE:  gogenx.GoCommand,
	}
)

func init() {

	goCmd.Flags().StringVar(&gogenx.VarStringDir, "dir", "", "The target dir")
	goCmd.Flags().StringVar(&gogenx.VarStringAPI, "api", "", "The api file")
	goCmd.Flags().StringVar(&gogenx.VarStringHome, "home", "", "The goctl home path of "+
		"the template, --home and --remote cannot be set at the same time, if they are, --remote "+
		"has higher priority")
	goCmd.Flags().StringVar(&gogenx.VarStringRemote, "remote", "", "The remote git repo "+
		"of the template, --home and --remote cannot be set at the same time, if they are, --remote"+
		" has higher priority\nThe git repo directory must be consistent with the "+
		"https://github.com/zeromicro/go-zero-template directory structure")
	goCmd.Flags().StringVar(&gogenx.VarStringBranch, "branch", "", "The branch of "+
		"the remote repo, it does work with --remote")
	goCmd.Flags().StringVar(&gogenx.VarStringStyle, "style", "gozero", "The file naming format,"+
		" see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]")

	rootCmd.AddCommand(goCmd)
}

// Execute executes the given command
func main() {
	os.Args = supportGoStdFlag(os.Args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}

func supportGoStdFlag(args []string) []string {
	copyArgs := append([]string(nil), args...)
	parentCmd, _, err := rootCmd.Traverse(args[:1])
	if err != nil { // ignore it to let cobra handle the error.
		return copyArgs
	}

	for idx, arg := range copyArgs[0:] {
		parentCmd, _, err = parentCmd.Traverse([]string{arg})
		if err != nil { // ignore it to let cobra handle the error.
			break
		}
		if !strings.HasPrefix(arg, dash) {
			continue
		}

		flagExpr := strings.TrimPrefix(arg, doubleDash)
		flagExpr = strings.TrimPrefix(flagExpr, dash)
		flagName, flagValue := flagExpr, ""
		assignIndex := strings.Index(flagExpr, assign)
		if assignIndex > 0 {
			flagName = flagExpr[:assignIndex]
			flagValue = flagExpr[assignIndex:]
		}

		if !isBuiltin(flagName) {
			// The method Flag can only match the user custom flags.
			f := parentCmd.Flag(flagName)
			if f == nil {
				continue
			}
			if f.Shorthand == flagName {
				continue
			}
		}

		goStyleFlag := doubleDash + flagName
		if assignIndex > 0 {
			goStyleFlag += flagValue
		}

		copyArgs[idx] = goStyleFlag
	}
	return copyArgs
}

func isBuiltin(name string) bool {
	return name == "version" || name == "help"
}
