package gowired

import (
	"reflect"

	"github.com/go-wired/models"
)

type Analizer struct {
	Output chan *models.Blueprint
}

func BuildAnalizer() *Analizer {
	return &Analizer{
		Output: make(chan *models.Blueprint),
	}
}

//Analize a object to createa  blueprint wit its dependencies.
func (a Analizer) Analize(itsSingleton bool, component interface{}) {
	//Get the name of the component
	componentType := reflect.TypeOf(component)

	//Start analizin and examine the componente type
	blueprint := a.examineType(componentType)

	blueprint.ItsSingleton = itsSingleton

}

//examineType check a type and extracts it dependencies, each type data get store on a blueprint
// and its emited through a channel @Output
func (a Analizer) examineType(componentType reflect.Type) *models.Blueprint {
	blueprint := &models.Blueprint{
		Type:         componentType,
		Name:         componentType.Name(),
		Dependencies: make([]models.FieldDep, componentType.NumField()),
	}
	//emit blueprint
	a.Output <- blueprint

	//Here we get dependenci information of this component
	for i := 0; i < componentType.NumField(); i++ {
		dependencyType := componentType.Field(i).Type

		if dependencyType.Kind() == reflect.Struct {
			blueprint.Dependencies = append(blueprint.Dependencies, models.FieldDep{
				Index: i,
				Name:  dependencyType.Name()})

			//do process again
			a.examineType(dependencyType)

		}
	}

	return blueprint
}
