package git

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/shared/resource/git/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformGitignore(m filesystem.Manager) error {
	component := component.MakeGitIgnore(m)

	if err := m.GenerateFile(component); err != nil {
		return err
	}

	output.ComponentCreated(component.Identifier)

	return nil
}
