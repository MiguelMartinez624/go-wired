package gowiredtest

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
