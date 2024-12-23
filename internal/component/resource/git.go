package resource

import (
	"github.com/charmingruby/bob/internal/component/resource/git/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformGitignore(m filesystem.Manager) {
	if err := m.GenerateFile(component.MakeGitIgnore(m)); err != nil {
		panic(err)
	}
}
