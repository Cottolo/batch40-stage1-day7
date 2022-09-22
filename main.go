package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	// route := mux.NewRouter()
	// route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// }).Methods("GET")

	route := mux.NewRouter()

	// path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	// routing
	route.HandleFunc("/home", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/detail-project/{id}", detailProject).Methods("GET")
	route.HandleFunc("/project", formAddProject).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")

	fmt.Println("server running at localhost:5000")
	http.ListenAndServe("localhost:5000", route)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/index.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	tmpl.Execute(w, nil)
}

func formAddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/project.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	tmpl.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	error := r.ParseForm()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Project Name :" + r.PostForm.Get("project-name"))
	fmt.Println("Start Date :" + r.PostForm.Get("start-date"))
	fmt.Println("End Date : " + r.PostForm.Get("end-date"))
	fmt.Println("Project Description :" + r.PostForm.Get("project-description"))
	fmt.Println("Technology : " + r.PostForm.Get("node-js"))
	fmt.Println("Technology : " + r.PostForm.Get("next-js"))
	fmt.Println("Technology : " + r.PostForm.Get("react-js"))
	fmt.Println("Technology : " + r.PostForm.Get("typescript"))

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/contact.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	tmpl.Execute(w, nil)
}

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/detail.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println(id)

	// OBJECT
	data := map[string]interface{}{
		"title":    "hello title",
		"containt": "containt",
		"id":       id,
	}

	tmpl.Execute(w, data)
}
