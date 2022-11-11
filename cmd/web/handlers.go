package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	blogs, err := app.infoBlogs.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.InfoBlogs = blogs
	println(data.InfoBlogs)
	app.render(w, http.StatusOK, "index.tmpl.html", data)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "about.tmpl.html", data)
}

func (app *application) blogView(w http.ResponseWriter, r *http.Request) {

}

func (app *application) blogPost(w http.ResponseWriter, r *http.Request) {

}
