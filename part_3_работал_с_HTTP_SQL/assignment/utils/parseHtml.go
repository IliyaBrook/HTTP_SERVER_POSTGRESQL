package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ParseHtml(w http.ResponseWriter, htmlStr string) {
	tmpl := template.Must(template.New("form").Parse(htmlStr))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Sprintln("Unable to execute html template:", err)
		log.Fatal(err)
	}
}
