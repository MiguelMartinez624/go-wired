package models

import "reflect"

//Blueprint should be a wrapper arround reflect type with fields and everything alrady for easy manipulation
//and extra data from the tags, some method to add fields and so on.
type Blueprint struct {
	Index        int
	Name         string
	ID           string
	SchemaID     string
	Childs       []*Blueprint
	ItsSingleton bool
}

//FieldDep each field of the object probably change name later
type FieldDep struct {
	Index int
	Name  string
	ID    string

	//Refelcts data to avoid use mehod each time its required
	//to know something about the object interface.
	Type reflect.Type
	Kind reflect.Kind
}
