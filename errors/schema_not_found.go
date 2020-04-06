package errors

import "fmt"

// SchemaNotFound schema not found
type SchemaNotFound struct {
	Selector string
}

func NewSchemaNotFound(selector string) *SchemaNotFound {
	return &SchemaNotFound{}
}

func (e SchemaNotFound) Error() string {
	return fmt.Sprintf("There is not a Schema for [%v].", e.Selector)
}
