package organism

import (
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/component/organism/module/component"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeAndRunBaseModule(m filesystem.Manager, module, database string) {
	newModule := component.MakeBaseRegistry(m, module)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	molecule.MakeAndRunCore(m, module, database)
	molecule.MakeAndRunRest(m, module)
}
