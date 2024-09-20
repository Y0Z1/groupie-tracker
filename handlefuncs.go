package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r,"/error-404", http.StatusSeeOther)
		return
	}
	t,err:=template.ParseFiles("templates/home.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ar := ParseHomeAPI()
	err = t.Execute(w,ar)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func artistpage(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/artist" {
		http.Redirect(w, r,"/error-404", http.StatusSeeOther)
		return
	}
	if r.Method!="POST"{
		http.Redirect(w, r,"/error-400", http.StatusSeeOther)
		return
	}
	id,_:=strconv.Atoi(r.FormValue("ID"))
	Year,_:=strconv.Atoi(r.FormValue("Year"))
	ar1:= Artist {
		Album: r.FormValue("Album"),
		ABM: r.FormValue("Albums"),
		Year: Year,
		ID: id,
		Name: r.FormValue("Name"),
		Memb: r.FormValue("Memb"),
		Memlen: r.FormValue("Memlen"),
		Image: r.FormValue("Image"),
		Locations: strings.Split(r.FormValue("Locations"),"/"),
		Rela: r.FormValue("Rela"),
	}
	t,err:=template.ParseFiles("templates/artist.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ar := ParseArtistAPI(ar1)
	err = t.Execute(w,ar)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func errorpage404(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/error-404" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("templates/error404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, "")
}

func errorpage400(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/error-400" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("templates/error400.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, "")
}