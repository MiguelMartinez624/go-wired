package errors

import "fmt"

//MissingImplementationTag missing implemts field
type MissingImplementationTag struct {
	Name string
}

func (e MissingImplementationTag) Error() string {
	return fmt.Sprintf("Struct [%v] its missing for implements field", e.Name)
}
