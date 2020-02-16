package gowired

import (
	"fmt"
	"reflect"

	"github.com/go-wired/models"
)

//Factory its the one who handles the creation of other objects///
type Factory struct {
	blueprints     *BlueprintMap
	interfacesImpl map[string]*models.Blueprint
	analizer       *Analizer
}

//CreateFactory create a factory this its the constructor function
//it initialize the bleprint map
func CreateFactory() *Factory {
	return &Factory{
		blueprints: NewBlueprintMap(),
		analizer:   BuildAnalizer(),
	}
}

func (f *Factory) RunFactory() {
	for blueprint := range f.analizer.Output {
		f.blueprints.AddBlueprint(blueprint)
		fmt.Println(f.blueprints.Length())

	}

}

// AddBlueprint register/add a blueprint to he factory so it can be use it later on to build
// the component
func (f *Factory) AddBlueprint(itsSingleton bool, component interface{}, name string) {

	f.analizer.Analize(component)
}

//CreateObjectByName create a object you can pass a name a object or anythin
func (f *Factory) CreateObjectByName(name interface{}) (obj interface{}) {
	blueprint, err := f.blueprints.FindBlueprint(name)
	if err != nil {
		panic(err)
	}
	//here we have the core object now we need to create its dependencies
	val, err := f.BuildObject(blueprint)

	for _, dep := range blueprint.Dependencies {
		if dep.Name == "" {
			continue
		}
		fmt.Println(dep)
		depBluep, err := f.blueprints.GetBlueprintByName(dep.Name)
		if err != nil {
			panic(err)
		}

		familyPtr := reflect.Indirect(val).Field(dep.Index)
		v := reflect.Indirect(familyPtr)
		fmt.Println(depBluep.Type)
		valDep := reflect.New(depBluep.Type)
		v.Set(reflect.Indirect(valDep))
		// .Set(reflect.New(depBluep.Type))
	}

	if err != nil {
		panic(err)

	}

	return val.Interface()
}

//BuildObject build and object using a blueprint
func (f *Factory) BuildObject(blueprint *models.Blueprint) (obj reflect.Value, err error) {
	value := reflect.New(blueprint.Type)

	return value, nil
}
