package gowired

import (
	"reflect"

	"github.com/go-wired/errors"
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
		f.blueprints[blueprint.Name] = blueprint
		// fmt.Printf("added %v \n", blueprint.Name)
	}
}

//FindBlueprint find a @Blueprint
func (f *BlueprintMap) FindBlueprint(identifier interface{}) (obj *models.Blueprint, err error) {
	kind := reflect.TypeOf(identifier).Kind()
	switch kind {
	case reflect.String:
		//TODO get blueprint by name directly
		name := identifier.(string)
		obj, err = f.GetBlueprintByName(name)
		return
	case reflect.Struct:
		name := reflect.TypeOf(identifier).Name()
		obj, err = f.GetBlueprintByName(name)
		return

	}
	return
}

func (f *BlueprintMap) GetBlueprintByName(name string) (bp *models.Blueprint, err error) {
	if blueprint, exist := f.blueprints[name]; exist {
		return blueprint, nil
	} else {
		err = errors.BlueprintNotFound{BlueprintName: name}
		return nil, err
	}
}

func (f *BlueprintMap) Length() int {
	return len(f.blueprints)
}
