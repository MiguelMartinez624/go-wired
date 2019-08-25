package gowired

import (
	"fmt"
	"go-wired/models"
	"reflect"
)

func getDependencies(obj interface{}) (deps []string, fields map[string]models.FieldDep) {
	deps = []string{}
	fields = map[string]models.FieldDep{}
	val := reflect.ValueOf(obj).Elem()
	mtype := val.Type()

	for i := 0; i < mtype.NumField(); i++ {
		field := mtype.Field(i).Type
		switch field.Kind() {
		case reflect.Interface:
			deps = append(deps, fmt.Sprintf("%v", mtype.Field(i).Type))
			fields[fmt.Sprintf("%v", mtype.Field(i).Type)] = models.FieldDep{
				index: i, name: mtype.Field(i).Name,
			}
		}
	}

	return
}
