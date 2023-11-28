package m2

import (
	"awesomeProject/pkg/component"
	"fmt"
	"github.com/defval/di"
)

type Config struct {
	di.Inject
	Param1 string `di:"cfg=m2.p1"`
	Param2 string `di:"cfg=m2.p2"`
}

type Module2 struct {
	Param1 string
	Param2 string
}

func NewModule2(cfg Config) *Module2 {
	fmt.Println("NewModule2")
	return &Module2{
		Param1: cfg.Param1,
		Param2: cfg.Param2,
	}
}

func (m Module2) ProvideDependencies(container *di.Container) error {
	container.ProvideValue(component.NewController("m2-controller1"))
	container.ProvideValue(component.NewController("m2-controller2"))
	return nil
}

func (m Module2) Start() {
	fmt.Printf("Module2 started with [%s, %s]\n", m.Param1, m.Param2)
}
