package gowired

import (
	"fmt"
	"go-wired/errors"
	"reflect"
)

//Factory its the one who handles the creation of other objects///
type Factory struct {
	Components     []interface{}
	store          map[string]*Element
	interfacesImpl map[string]*Element
}

//Init should init the node
func (factory *Factory) Init() {
	//just register the components on the slice of Components
	for i := 0; i < len(factory.Components); i++ {
		factory.AddBlueprint(factory.Components[i])
	}

	for _, p := range factory.store {

		for i := 0; i < len(p.dependencies); i++ {

			dep := p.dependencies[i]

			if storedDep, have := factory.interfacesImpl[dep]; have == true {
				switch p.Kind {
				case reflect.Struct:
					fieldToChange := p.Value.FieldByName(p.fieldDep[dep].name)
					fieldToChange.Set(storedDep.Value.Addr())
				}
			}
		}
	}
}

//GetComponent return a component instnace with all dependencies injected(singleton)
func (factory *Factory) GetComponent(descripcion interface{}) interface{} {
	//get the type of the component that to return.
	typeDesc := reflect.TypeOf(descripcion)
	fmt.Printf("Getting value of %v \n", typeDesc.Name())
	fmt.Printf("%v \n", factory.store)
	//search for the component on the store if the component exist it will be returned
	// otherwise will return nil
	if comp, have := factory.store[typeDesc.Name()]; have {
		return comp.Element
	}

	return nil
}

//AddBlueprint register a obj as a element of the node
func (factory *Factory) AddBlueprint(obj interface{}) (ele *Element, err error) {
	//Check that the store its initialized before attempting to store a blueprint
	if factory.store == nil {
		factory.store = make(map[string]*Element)
		factory.interfacesImpl = make(map[string]*Element)
	}

	val := reflect.ValueOf(obj).Elem()
	typeOfT := val.Type()

	//element that have the tag of implements
	field, itImplementField := typeOfT.FieldByName("implements")
	if itImplementField == false {
		return nil, errors.MissingImplementationTag{Name: typeOfT.Name()}
	}

	deps, fields := getDependencies(obj)

	ele = &Element{
		Name:    typeOfT.Name(),
		Element: obj,
		interfaces: []string{
			field.Tag.Get("implements"),
		},
		Type:         val.Type(),
		Value:        val,
		Kind:         val.Kind(),
		fieldDep:     fields,
		dependencies: deps,
	}

	factory.store[typeOfT.Name()] = ele
	factory.interfacesImpl[field.Tag.Get("implements")] = ele

	fmt.Printf("Storing VAL: %v \n", factory.store)
	return ele, nil
}
