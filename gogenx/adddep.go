package gogenx

import (
	"errors"
	"os"
	"strings"

	"github.com/gioco-play/goctl-gfdep"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/util"
)

var (
	// VarStringName describes the dependency name to inject via add-dep.
	VarStringName string

	errMissingDepName = errors.New("add-dep: missing --name")
)

// AddDep injects a dependency (declared as <tplDir>/deps/<name>.tpl in the
// --home or --remote template repo) into the service rooted at the current
// working directory.
func AddDep(_ *cobra.Command, _ []string) error {
	name := VarStringName
	if len(name) == 0 {
		return errMissingDepName
	}

	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	if len(remote) > 0 {
		repo, err := util.CloneIntoGitHome(remote, branch)
		if err != nil {
			return err
		}
		home = repo
	}

	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, n := range strings.Split(name, ",") {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		if err := dep.Inject(workDir, home, n); err != nil {
			return err
		}
	}
	return nil
}
