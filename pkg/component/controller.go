package component

import "fmt"

type Controller struct {
	Name string
}

func NewController(name string) Controller {
	return Controller{
		Name: name,
	}
}

func (c Controller) Start() {
	fmt.Printf("Controller %s started\n", c.Name)
}
