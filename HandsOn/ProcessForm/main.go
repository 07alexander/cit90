package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

type person struct {
	FirstName string
	LastName string
	Subscribed bool
}

func main(){
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, req *http.Request){
	f:=req.FormValue("first")
	l:=req.FormValue("last")
	s:=req.FormValue("subscribe") == "on"

	err:= tpl.ExecuteTemplate(w, "index.gohtml", person{f,l,s})
	if err!=nil{
		http.Error(w,err.Error(),500)
		log.Fatalln(err)
	}

}