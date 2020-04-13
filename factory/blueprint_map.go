package factory

import (
	"reflect"
)

//BlueprintMap its the store for blueprint
type BlueprintMap struct {
	blueprints map[string]*Blueprint
}

//NewBlueprintMap contructor for Blueprint
func NewBlueprintMap() *BlueprintMap {
	return &BlueprintMap{
		blueprints: make(map[string]*Blueprint, 10),
	}
}

func (f *BlueprintMap) ListenStream(in chan *Blueprint) {
	for blueprint := range in {
		f.AddBlueprint(blueprint)
	}
}

//AddBlueprint register a blueprint to the list
func (f *BlueprintMap) AddBlueprint(blueprint *Blueprint) {
	if _, exist := f.blueprints[blueprint.ID]; !exist {
		//Get the type of the element to store in the blueprint
		f.blueprints[blueprint.ID] = blueprint
	}
}

//FindBlueprint find a @Blueprint
func (f *BlueprintMap) FindBlueprint(identifier interface{}) (obj *Blueprint, err error) {
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

func (f *BlueprintMap) GetBlueprintByName(name string) (bp *Blueprint, err error) {
	if blueprint, exist := f.blueprints[name]; exist {
		return blueprint, nil
	} else {
		err = BlueprintNotFound{BlueprintName: name}
		return nil, err
	}
}

func (f *BlueprintMap) Length() int {
	return len(f.blueprints)
}
