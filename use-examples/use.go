package main

import (
	gowired "github.com/go-wired"
)

type ComponentOne struct {
	Name string
	ID   string
}
type ComponentTwo struct {
	Name          string
	ID            string
	dependencyOne ComponentOne
}

func main() {
	factory := gowired.CreateFactory()

	factory.AddBlueprint(true, ComponentOne{}, "ComponentOne")
	factory.AddBlueprint(true, ComponentTwo{}, "ComponentTwo")
	// factory.CreateObjectByName(ComponentOne{})
}
