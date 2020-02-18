package gowired

import (
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
	factory := &Factory{
		blueprints: NewBlueprintMap(),
		analizer:   BuildAnalizer(),
	}
	go factory.RunFactory()
	return factory
}

//RunFactory init all the process that the factory need for work
// 1- Start listining for events from the analizer so each time a blueprint
//    its created it will make sure that its property stored on the map for
//    later use
func (f *Factory) RunFactory() {

	//listen to incoming blueprints
	for blueprint := range f.analizer.Output {
		f.blueprints.AddBlueprint(blueprint)
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
	prtVal, err := f.BuildObject(blueprint)
	if err != nil {
		panic(err)

	}
	//Set all dependencies of the object.
	f.setDependencies(prtVal, blueprint)

	return prtVal.Interface()
}

//setDependencies iterates though all @FieldDep and using the @Blueprints stored in
// the map it will procced to create and assign the value field, this its done on a
// recursive manner so it guarantee the childs tree dependencies are fullfilled too
func (f *Factory) setDependencies(prtVal reflect.Value, blueprint *models.Blueprint) {
	//indiect the value of the Ptr to be able to work fields
	val := reflect.Indirect(prtVal)
	//For each dependencie on the @Blueprint it will get the blueprint dependency
	//build and object and assing it to the correspondent field.
	//
	for _, dep := range blueprint.Dependencies {
		//On the analize process some space are left on the array of dependencies
		//this validate temporally that the index it have a valid @FieldDep value.
		if dep.Name == "" {
			continue
		}
		depBluep, err := f.blueprints.GetBlueprintByName(dep.Name)
		if err != nil {
			panic(err)
		}

		//this its the dependency object instance pointer
		valDepPtr := reflect.New(depBluep.Type)
		f.setDependencies(valDepPtr, depBluep)
		val.Field(dep.Index).Set(reflect.Indirect(valDepPtr))

	}
}

//BuildObject build and object using a blueprint
func (f *Factory) BuildObject(blueprint *models.Blueprint) (obj reflect.Value, err error) {
	value := reflect.New(blueprint.Type)

	return value, nil
}
