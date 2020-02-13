package gowired

import (
	"fmt"
	"github.com/go-wired/models"
	"reflect"
)

func Analize(component interface{}) (blueprint *models.Blueprint) {
	t := reflect.TypeOf(component)
	// v := reflect.ValueOf(component)
	blueprint = &models.Blueprint{}
	for i := 0; i < t.NumField(); i++ {
		//name of the field
		// dependencyName := t.Field(i).Name
		//this tis the type of te dependency
		blueprint.Type = t.Field(i).Type

	}
	fmt.Printf("Blueprint:  %v\n", blueprint)

	return blueprint
}
