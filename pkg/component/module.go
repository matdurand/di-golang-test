package component

import "github.com/defval/di"

type Module interface {
	ProvideDependencies(container *di.Container) error
	Start()
}
