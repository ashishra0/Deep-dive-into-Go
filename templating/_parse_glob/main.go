package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template // assigned for global scope. So this is accessible throughout the package.

func init() { // This function runs once when the program is starting up.
	tpl = template.Must(template.ParseGlob("templates/*")) // template.Must will do the error checking on the pointer to template and return the pointer to template
}

func main() {

	players := map[string]string{
		"India":   "Chhetri",
		"America": "Pulisic",
		"Spain":   "Ceballos",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "three.txt", players)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
