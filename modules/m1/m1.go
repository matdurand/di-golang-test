package m1

import (
	"awesomeProject/pkg/component"
	"fmt"
	"github.com/defval/di"
)

type Config struct {
	di.Inject
	Value1 string `di:"cfg=m1.v1"`
	Value2 string `di:"cfg=m1.v2"`
}

type Module1 struct {
	Value1 string
	Value2 string
}

func NewModule1(cfg Config) *Module1 {
	fmt.Println("NewModule1")
	return &Module1{
		Value1: cfg.Value1,
		Value2: cfg.Value2,
	}
}

func (m Module1) ProvideDependencies(container *di.Container) error {
	container.ProvideValue(component.NewController("m1-controller1"))
	container.ProvideValue(component.NewController("m1-controller2"))
	return nil
}

func (m Module1) Start() {
	fmt.Printf("Module1 started with [%s, %s]\n", m.Value1, m.Value2)
}
