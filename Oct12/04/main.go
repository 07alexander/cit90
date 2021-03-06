package main

import (
	"net/http"
	"log"
	"html/template"
)
var tpl *template.Template
func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/about",about)
	http.HandleFunc("/contact",contact)
	http.HandleFunc("/apply",apply)
	http.HandleFunc("/applyProcess",applyProcess)
	http.ListenAndServe(":8080", nil)
}

func init(){
	tpl=template.Must(template.ParseGlob("templates/*"))
}


func index(w http.ResponseWriter, _*http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, _*http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, _*http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, _*http.Request) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}