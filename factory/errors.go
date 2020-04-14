package factory

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

//BlueprintNotFound missing implemts field
type BlueprintNotFound struct {
	BlueprintName string
}

func (e BlueprintNotFound) Error() string {
	return fmt.Sprintf("There is not a Blueprint for [%v] registered on the factory.", e.BlueprintName)
}

// ProviderNotFound schema not found
type ProviderNotFound struct {
	Selector string
}

func NewProviderNotFound(selector string) *ProviderNotFound {
	return &ProviderNotFound{Selector: selector}
}

func (e ProviderNotFound) Error() string {
	return fmt.Sprintf("There is not a provider for [%v].", e.Selector)
}
