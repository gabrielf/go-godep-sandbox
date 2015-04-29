package main

import "github.com/gabrielf/godep-sandbox/Godeps/_workspace/src/github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
