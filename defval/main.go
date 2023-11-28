package main

import (
	"awesomeProject/modules/m1"
	"awesomeProject/modules/m2"
	"awesomeProject/pkg/component"
	"github.com/defval/di"
)

func startRoutes(controllers []component.Controller) {
	for _, c := range controllers {
		c.Start()
	}
}

func startModules(rootDI *di.Container, modules []component.Module) error {
	var err error
	for _, m := range modules {
		moduleDI, _ := di.New()
		err = moduleDI.AddParent(rootDI)
		if err != nil {
			return err
		}
		err = m.ProvideDependencies(moduleDI)
		if err != nil {
			return err
		}
		err = moduleDI.Invoke(startRoutes)
		if err != nil {
			return err
		}
		m.Start()
	}
	return nil
}

func main() {
	c, err := di.New(
		di.ProvideValue("aaa", di.Tags{"cfg": "m1.v1"}),
		di.ProvideValue("bbb", di.Tags{"cfg": "m1.v2"}),
		di.ProvideValue("ccc", di.Tags{"cfg": "m2.p1"}),
		di.ProvideValue("ddd", di.Tags{"cfg": "m2.p2"}),
		di.Provide(m1.NewModule1, di.As(new(component.Module))),
		di.Provide(m2.NewModule2, di.As(new(component.Module))),
	)
	if err != nil {
		panic(err)
	}

	err = c.Invoke(startModules)
	if err != nil {
		panic(err)
	}
}
