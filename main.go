package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

type Person struct {
	Name string
	Age  string
}

func main() {
	a := Person{
		Name: "Richa",
		Age:  "27",
	}
	b := Person{
		Name: "Sahil",
		Age:  "27",
	}
	persons := []Person{a, b}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", persons)
	if err != nil {
		log.Fatal(err)
	}
}
