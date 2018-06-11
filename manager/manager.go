package manager

type Manager interface {
}

type DefaultManager struct {
}

func NewManager() Manager {
	return &DefaultManager{}
}
