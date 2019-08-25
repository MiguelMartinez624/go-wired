package gowired

import (
	"fmt"
	"reflect"
)

//Element should be a wrapper arround reflect type with fields and everything alrady for easy manipulation
//and extra data from the tags, some method to add fields and so on.
type Element struct {
	Name         string
	Element      interface{}
	fieldDep     map[string]FieldDep
	interfaces   []string
	dependencies []string
	Type         reflect.Type
	Kind         reflect.Kind
	Value        reflect.Value
}

//FieldDep each field of the object probably change name later
type FieldDep struct {
	index int
	name  string
}

//Builder its a piece of funcionality, all structs (obj) related to that unique feature, related to a domain or
//related domains, all assets should be contained into a Nodo
type Builder struct {
	Components     []interface{}
	store          map[string]Element
	interfacesImpl map[string]Element
}



func getDependencies(obj interface{}) (deps []string, fields map[string]FieldDep) {
	deps = []string{}
	fields = map[string]FieldDep{}
	val := reflect.ValueOf(obj).Elem()
	mtype := val.Type()

	for i := 0; i < mtype.NumField(); i++ {
		field := mtype.Field(i).Type
		switch field.Kind() {
		case reflect.Interface:
			deps = append(deps, fmt.Sprintf("%v", mtype.Field(i).Type))
			fields[fmt.Sprintf("%v", mtype.Field(i).Type)] = FieldDep{
				index: i, name: mtype.Field(i).Name,
			}
		}
	}

	return
}
