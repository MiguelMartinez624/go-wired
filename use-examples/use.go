package main

import (
	factory "github.com/miguelmartinez624/go-wired"
	factorytest "github.com/miguelmartinez624/go-wired/tests"
)

func main() {
	for i := 0; i < 1000; i++ {
		f := factory.CreateFactory()
		f.RegisterObject(factorytest.ComponentOne{})

		f.RegisterProvider("github.com/miguelmartinez624/go-wired/tests.Dummer", factorytest.BasicDummer{})

		c := f.CreateObjectByName(factorytest.ComponentOne{}).(*factorytest.ComponentOne)
		c.DrummerImpl.Dumb()
	}
}
