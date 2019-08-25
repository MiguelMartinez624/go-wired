package gowired

import (
	"fmt"
	"reflect"
	"go-wired/errors"
	"go-wired/models"
)

//Factory its the one who handles the creation of other objects///
type Factory struct {
	Components     []interface{}
	store          map[string]*models.Blueprint
	interfacesImpl map[string]*models.Blueprint
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

	//search for the component on the store if the component exist it will be returned
	// otherwise will return nil
	if comp, have := factory.store[typeDesc.Name()]; have {
		return comp.models.Blueprint
	}

	return nil
}

//AddBlueprint register a obj as a models.Blueprint of the node
func (factory *Factory) AddBlueprint(obj interface{}) (ele *models.Blueprint, err error) {
	//Check that the store its initialized before attempting to store a Blueprint
	if factory.store == nil {
		factory.store = make(map[string]*models.Blueprint)
		factory.interfacesImpl = make(map[string]*models.Blueprint)
	}

	val := reflect.ValueOf(obj).Elem()
	typeOfT := val.Type()

	//models.Blueprint that have the tag of implements
	field, itImplementField := typeOfT.FieldByName("implements")
	if itImplementField == false {
		return nil, errors.MissingImplementationTag{Name: typeOfT.Name()}
	}

	deps, fields := getDependencies(obj)

	ele = &models.Blueprint{
		Name:             typeOfT.Name(),
		models.Blueprint: obj,
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
