package gowiredtest

import (
	"testing"

	gowired "github.com/go-wired"
)

type TestFunction = func(f *gowired.Factory, t *testing.T)

type TestCase struct {
	Name     string
	CaseFunc TestFunction
}

func Suite(f *gowired.Factory, t *testing.T) {
	ttcc := []TestCase{
		{Name: "Flat Dependency", CaseFunc: FlatDependencyTest},
		{Name: "Nested Dependency", CaseFunc: NestedDependencyTest},
	}

	for _, tc := range ttcc {
		t.Run(tc.Name, func(t *testing.T) { tc.CaseFunc(f, t) })
	}
}

//FlatDependencyTest 1 level dependency test
func FlatDependencyTest(f *gowired.Factory, t *testing.T) {

	f.AddBlueprint(false, ComponentOne{}, "")
	componentOne := f.CreateObjectByName(ComponentOne{}).(*ComponentOne)
	depRef := &componentOne.NodeOne

	if componentOne == nil {
		t.Error("Creation fail")
	}

	if depRef == nil {
		t.Error("Dependency failed")
	}
}

//NestedDependencyTest 2 level dependency test
func NestedDependencyTest(f *gowired.Factory, t *testing.T) {

	f.AddBlueprint(false, ComponentTwo{}, "")
	componentTwo := f.CreateObjectByName(ComponentTwo{}).(*ComponentTwo)

	if componentTwo == nil {
		t.Error("Creation fail")
	}

	depRef1 := &componentTwo.DependencyOne
	if depRef1 == nil {
		t.Error("Dependency failed")
	}

	depRef2 := &componentTwo.DependencyOne.NodeOne
	if depRef2 == nil {
		t.Error("Dependency failed")
	}
}

func TestFacory(t *testing.T) {
	factory := gowired.CreateFactory()

	Suite(factory, t)
}
