package main

import (
	gowired "github.com/miguelmartinez624/go-wired"
	gowiredtest "github.com/miguelmartinez624/go-wired/tests"
)

func main() {
	for i := 0; i < 1000; i++ {
		f := gowired.CreateFactory()
		f.RegisterObject(gowiredtest.ComponentOne{})

		f.RegisterProvider("github.com/miguelmartinez624/go-wired/tests.Dummer", gowiredtest.BasicDummer{})

		c := f.CreateObjectByName(gowiredtest.ComponentOne{}).(*gowiredtest.ComponentOne)
		c.DrummerImpl.Dumb()
	}
}
