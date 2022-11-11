package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Initialize a slice containing the paths to the two files. It's important // to note that the file containing our base template must be the *first* // file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/index.tmpl",
	}
	// Use the template.ParseFiles() function to read the files and store the // templates in a template set. Notice that we can pass the slice of file // paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base" // template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/infoBlog/about" {
		http.NotFound(w, r)
		return
	}
	// Initialize a slice containing the paths to the two files. It's important // to note that the file containing our base template must be the *first* // file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/about.tmpl",
	}
	// Use the template.ParseFiles() function to read the files and store the // templates in a template set. Notice that we can pass the slice of file // paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base" // template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func samplePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/infoBlog/samplePost" {
		http.NotFound(w, r)
		return
	}
	// Initialize a slice containing the paths to the two files. It's important // to note that the file containing our base template must be the *first* // file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/samplePost.tmpl",
	}
	// Use the template.ParseFiles() function to read the files and store the // templates in a template set. Notice that we can pass the slice of file // paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base" // template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/infoBlog/contact" {
		http.NotFound(w, r)
		return
	}
	// Initialize a slice containing the paths to the two files. It's important // to note that the file containing our base template must be the *first* // file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/contact.tmpl",
	}
	// Use the template.ParseFiles() function to read the files and store the // templates in a template set. Notice that we can pass the slice of file // paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base" // template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/infoBlog/login" {
		http.NotFound(w, r)
		return
	}
	// Initialize a slice containing the paths to the two files. It's important // to note that the file containing our base template must be the *first* // file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/login.tmpl",
	}
	// Use the template.ParseFiles() function to read the files and store the // templates in a template set. Notice that we can pass the slice of file // paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base" // template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
