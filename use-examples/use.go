package main

import (
	"fmt"

	factory "github.com/miguelmartinez624/go-wired/factory"
	factorytest "github.com/miguelmartinez624/go-wired/factory/tests"
)

type InstanceLoca struct {
	Name string
}

func (i InstanceLoca) Dumb() {
	fmt.Print("LOOOOOOOOOOOL")
}

func main() {
	u := InstanceLoca{Name: "LE ORIGINAL"}
	// for i := 0; i < 1000; i++ {
	f := factory.CreateFactory()
	f.GenerateObjectSchema(factorytest.ComponentOne{})

	f.RegisterProviderInstance("github.com/miguelmartinez624/go-wired/factory/tests.Dummer", u)

	c := f.CreateObjectByName(factorytest.ComponentOne{}).(*factorytest.ComponentOne)
	// }
	fmt.Println(c.DrummerImpl)
}
