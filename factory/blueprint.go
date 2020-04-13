package factory

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
