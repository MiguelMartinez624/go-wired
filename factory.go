package gowired

import (
	"github.com/go-wired/models"
)

//Factory its the one who handles the creation of other objects///
type Factory struct {
	blueprints     *BlueprintMap
	interfacesImpl map[string]*models.Blueprint
}

//CreateFactory create a factory this its the constructor function
//it initialize the bleprint map
func CreateFactory() *Factory {
	return &Factory{
		blueprints: NewBlueprintMap(),
	}
}

// AddBlueprint register/add a blueprint to he factory so it can be use it later on to build
// the component
func (f *Factory) AddBlueprint(itsSingleton bool, component interface{}, name string) {
	Analize(component)
	f.blueprints.AddBlueprint(
		&models.Blueprint{ItsSingleton: itsSingleton, Name: name, Element: component})
}

//CreateObjectByName create a object you can pass a name a object or anythin
func (f *Factory) CreateObjectByName(name interface{}) (obj interface{}, err error) {
	blueprint, err := f.blueprints.FindBlueprint(name)
	if err != nil {
		return nil, err
	}

	obj, err = f.BuildObject(blueprint)
	if err != nil {
		return nil, err
	}
	return
}

//BuildObject build and object using a blueprint
func (f *Factory) BuildObject(blueprint *models.Blueprint) (obj interface{}, err error) {
	return
}
