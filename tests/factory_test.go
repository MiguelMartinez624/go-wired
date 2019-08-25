package test

import (
	"go-wired"
	"testing"
)

// FactorySuite uses cases of factory object.
func FactorySuite(t *testing.T, factory *gowired.Factory) {

	tt := []struct {
		name string
		cb   func(t *testing.T, factory *gowired.Factory)
	}{
		{name: "register component", cb: RegisterComponent},
		{name: "register component without implement field", cb: RegisterComponentWithoutImplementField},
		{name: "get componente registed", cb: GetComponentSucced},
		{name: "get componente unregisted", cb: GetComponentFail},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.cb(t, factory)
		})
	}
}

// RegisterComponent test the registration of a component into the factory
func RegisterComponent(t *testing.T, factory *gowired.Factory) {

	_, err := factory.AddBlueprint(&Struct1{})

	if err != nil {
		t.Errorf(err.Error())
		return
	}

}

// RegisterComponent test the registration of a component into the factory
func RegisterComponentWithoutImplementField(t *testing.T, factory *gowired.Factory) {
	elementInfo, _ := factory.AddBlueprint(&Struct2{})

	if elementInfo != nil {
		t.Errorf("MissingImplementationTag expected as result")

	}
}

// GetComponentSucced test the registration of a component into the factory
func GetComponentSucced(t *testing.T, factory *gowired.Factory) {
	component := factory.GetComponent(Struct1{}).(*Struct1)

	if component == nil {
		t.Errorf("Component spected")
	}

}

// GetComponentFail test the registration of a component into the factory
func GetComponentFail(t *testing.T, factory *gowired.Factory) {
	component := factory.GetComponent(Struct3{})
	if component != nil {
		t.Errorf("Should return nil for unregistered components")
	}
}

// TestFactory test basic funcionality of a Factory object
// the main class of the package
func TestFactory(t *testing.T) {
	factory := &gowired.Factory{}
	FactorySuite(t, factory)
}
