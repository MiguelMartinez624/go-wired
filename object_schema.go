package gowired

import "reflect"

// ObjectSchema
type ObjectSchema struct {
	ID        string
	Name      string
	Package   string
	Raw       interface{}
	FieldsMap map[int]*ObjectSchema
	//Refelcts data to avoid use mehod each time its required
	//to know something about the object interface.
	Type  reflect.Type
	Kind  reflect.Kind
	Value reflect.Value
}
