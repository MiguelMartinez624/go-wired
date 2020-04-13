package factorytest

import "fmt"

type Dummer interface {
	Dumb()
}

type BasicDummer struct {
}

func (d BasicDummer) Dumb() {
	fmt.Println("LOL")
}

type GrandChild struct {
	Name string
	ID   string
}

type ComponentOne struct {
	Name        string
	NodeOne     GrandChild
	DrummerImpl Dummer
}

type ComponentTwo struct {
	Name          string
	ID            string
	DependencyOne ComponentOne
}

type ComponentThree struct {
	Name string
	ID   string
}
