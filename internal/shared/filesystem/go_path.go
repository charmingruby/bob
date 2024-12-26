package filesystem

import "fmt"

func (m *Manager) RootPath() string {
	return m.Data
}

func (m *Manager) DependencyPath() string {
	return fmt.Sprintf("%s/%s", m.Data, m.SourceDirectory)
}
