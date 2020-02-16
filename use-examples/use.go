package main

import (
	"fmt"
	gowired "github.com/go-wired"
)

type ComponentOne struct {
	Name string
	ID   string
}
type ComponentTwo struct {
	Name          string
	ID            string
	DependencyOne ComponentOne
}

func main() {
	factory := gowired.CreateFactory()

	go factory.RunFactory()
	factory.AddBlueprint(true, ComponentTwo{}, "ComponentTwo")

	componentOne := factory.CreateObjectByName(ComponentTwo{}).(*ComponentTwo)
	componentOne.DependencyOne.Name = "lolazo co'o e tu madre"
	fmt.Println(componentOne)
}
