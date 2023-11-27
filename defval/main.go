package main

import (
	"awesomeProject/modules/m1"
	"awesomeProject/modules/m2"
	"fmt"
	"github.com/defval/di"
)

func doM1(m1 m1.Module1) {
	m1.Start()
	//fmt.Println("do m1")
}

func doM2(m2 m2.Module2) {
	m2.Start()
	//fmt.Println("do m2")
}

func main() {
	c, err := di.New(
		di.ProvideValue("aaa", di.Tags{"cfg": "m1.v1"}),
		di.ProvideValue("bbb", di.Tags{"cfg": "m1.v2"}),
		di.ProvideValue("ccc", di.Tags{"cfg": "m2.p1"}),
		di.ProvideValue("ddd", di.Tags{"cfg": "m2.p2"}),
		di.Provide(m1.NewModule1),
		di.Provide(m2.NewModule2),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("start...")
	err = c.Invoke(doM1)
	err = c.Invoke(doM1)
	if err != nil {
		panic(err)
	}
	err = c.Invoke(doM2)
	if err != nil {
		panic(err)
	}
}
