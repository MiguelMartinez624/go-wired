package main

import (
	gowired "github.com/go-wired"
	gowiredtest "github.com/go-wired/tests"
)

func main() {
	for i := 0; i < 100; i++ {
		f := gowired.CreateFactory()
		f.RegisterObject(gowiredtest.ComponentOne{})

		f.RegisterProvider("github.com/go-wired/tests.Dummer", gowiredtest.BasicDummer{})

		c := f.CreateObjectByName(gowiredtest.ComponentOne{}).(*gowiredtest.ComponentOne)
		c.DrummerImpl.Dumb()
	}
}
