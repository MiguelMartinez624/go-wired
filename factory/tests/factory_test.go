package factorytest

import (
	"testing"

	"github.com/miguelmartinez624/go-wired/factory"
)

type TestFunction = func(f *factory.Factory, t *testing.T)

type TestCase struct {
	Name     string
	CaseFunc TestFunction
}

func Suite(f *factory.Factory, t *testing.T) {
	ttcc := []TestCase{}

	for _, tc := range ttcc {
		t.Run(tc.Name, func(t *testing.T) { tc.CaseFunc(f, t) })
	}
}

func TestFacory(t *testing.T) {
	factory := factory.CreateFactory()

	Suite(factory, t)
}
