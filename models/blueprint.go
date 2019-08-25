package models

// Blueprint represent a object spec, like interfaces implemented.
type Blueprint struct {
	//interfaces implemented by this blueprint should be under a tag
	// with the package name
	interfaces []string
}
