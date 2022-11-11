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

	app.render(w, http.StatusOK, "index.html", data)
}

func (app *application) blogView(w http.ResponseWriter, r http.Request) {

}

func (app *application) blogPost(w http.ResponseWriter, r http.Request) {

}
