package main

import (
	"fmt"
	gowired "github.com/go-wired"
)

type GrandChild struct {
	Name string
	ID   string
}

type ComponentOne struct {
	Name    string
	NodeOne GrandChild
}
type ComponentTwo struct {
	Name          string
	ID            string
	DependencyOne ComponentOne
}

func main() {
	factory := gowired.CreateFactory()

	factory.AddBlueprint(true, ComponentTwo{}, "ComponentTwo")

	componentOne := factory.CreateObjectByName(ComponentTwo{}).(*ComponentTwo)
	componentOne.DependencyOne.Name = "lolazo co'o e tu madre"
	componentOne.DependencyOne.NodeOne.Name = "YO soy el maldito abluelo"
	fmt.Printf("%+v\n", componentOne)
}
