package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("file.gohtml") // template.ParseFiles returns a pointer to template and an error object.
	// So after that we use pointer to template to perform the Exedcute method which is a recevier function to pointer to template
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
