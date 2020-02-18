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

func (a Analizer) Analize(component interface{}) {
	componentType := reflect.TypeOf(component)
	//Get the name of the component
	blueprint := a.examineType(componentType)

	a.Output <- blueprint
}

func (a Analizer) examineType(componentType reflect.Type) *models.Blueprint {
	blueprint := &models.Blueprint{
		Type:         componentType,
		Name:         componentType.Name(),
		Dependencies: make([]models.FieldDep, componentType.NumField()),
	}

	//Here we get dependenci information of this component
	for i := 0; i < componentType.NumField(); i++ {
		dependencyType := componentType.Field(i).Type
		if dependencyType.Kind() == reflect.Struct {
			blueprint.Dependencies = append(blueprint.Dependencies, models.FieldDep{
				Index: i,
				Name:  dependencyType.Name()})
			//do process again
			depBlueprint := a.examineType(dependencyType)
			a.Output <- depBlueprint

		}

	}
	return blueprint
}
