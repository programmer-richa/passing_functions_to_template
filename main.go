package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc":    strings.ToUpper,
	"ft":    firstThree,
	"ftime": formatTime,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func formatTime(t time.Time) string {

	return t.Format("Jan 1, 2020")
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

	err = tpl.ExecuteTemplate(os.Stdout, "time.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
