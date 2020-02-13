package gowired

import (
	"fmt"
	"reflect"

	"github.com/go-wired/models"
)

//BlueprintMap its the store for blueprint
type BlueprintMap struct {
	blueprints map[string]*models.Blueprint
}

//NewBlueprintMap contructor for Blueprint
func NewBlueprintMap() *BlueprintMap {
	return &BlueprintMap{
		blueprints: make(map[string]*models.Blueprint, 10),
	}
}

//AddBlueprint register a blueprint to the list
func (f *BlueprintMap) AddBlueprint(blueprint *models.Blueprint) {
	if _, exist := f.blueprints[blueprint.Name]; !exist {
		//Get the type of the element to store in the blueprint
		elementType := reflect.TypeOf(blueprint.Element)
		//Get the name of the component
		blueprint.Name = elementType.Name()

		f.blueprints[blueprint.Name] = blueprint
	}
}

//FindBlueprint find a @Blueprint
func (f *BlueprintMap) FindBlueprint(identifier interface{}) (obj *models.Blueprint, err error) {
	kind := reflect.TypeOf(identifier).Kind()
	switch kind {
	case reflect.String:
		//TODO get blueprint by name directly
		// if _, exist := f.blueprints[name.(string)]; exist {
		// 	//Use the @Blueprint
		// }
		break
	case reflect.Struct:
		//TODO get blueprint by getting the name from the @reflect.Type
		fmt.Println(kind)

		break

	}
	return
}
