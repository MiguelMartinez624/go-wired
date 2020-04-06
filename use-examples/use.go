package main

import (
	gowired "github.com/go-wired"
	gowiredtest "github.com/go-wired/tests"
)

func main() {

	f := gowired.CreateFactory()

	schemaID := f.GenerateObjectSchema(gowiredtest.ComponentOne{})
	_, err := f.CreateBlueprint(schemaID)
	if err != nil {
		panic(err)
	}

	f.RegisterProvider("github.com/go-wired/tests.Dummer", gowiredtest.BasicDummer{})

	c := f.CreateObjectByName(gowiredtest.ComponentOne{}).(*gowiredtest.ComponentOne)
	c.DrummerImpl.Dumb()

}
