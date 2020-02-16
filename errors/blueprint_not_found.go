package errors

import "fmt"

//BlueprintNotFound missing implemts field
type BlueprintNotFound struct {
	BlueprintName string
}

func (e BlueprintNotFound) Error() string {
	return fmt.Sprintf("There is not a Blueprint for [%v] registered on the factory.", e.BlueprintName)
}
