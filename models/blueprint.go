package models

import "reflect"

//Blueprint should be a wrapper arround reflect type with fields and everything alrady for easy manipulation
//and extra data from the tags, some method to add fields and so on.
type Blueprint struct {
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
