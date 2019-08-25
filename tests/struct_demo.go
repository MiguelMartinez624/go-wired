package test

//Interf1 test interface
type Interf1 interface {
	Hello() string
}

// Struct1 demo
type Struct1 struct {
	implements string `implements:"test.Interf1"`
}

//Hello says hello
func (s *Struct1) Hello() string {
	return "Hello"
}

//Struct2  demo
type Struct2 struct {
}

//Struct3 demo
type Struct3 struct {
}
