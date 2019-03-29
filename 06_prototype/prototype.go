package _6_prototype

type Cloneable interface {
	Clone() Cloneable
}
type Manager struct {
	container map[string]Cloneable
}

func NewManager() Manager {
	return Manager{
		make(map[string]Cloneable),
	}
}

func (m *Manager) Get(key string) Cloneable {
	return m.container[key]
}

func (m *Manager) Set(key string, prototype Cloneable) {
	m.container[key] = prototype
}
