package main

import (
	"github.com/go-martini/martini"
	"github.com/zongh1314/yaag/martiniyaag"
	"github.com/zongh1314/yaag/yaag"
)

func main() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Martini", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})
	m := martini.Classic()
	m.Use(martiniyaag.Document)
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
