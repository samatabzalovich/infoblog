package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r http.Request) {
	blogs, err := app.infoBlogs.Insert()
	if err != nil {
		app.serverError(v, err)
		return
	}
}

func (app *application) blogView(w http.ResponseWriter, r http.Request) {

}

func (app *application) blogPost(w http.ResponseWriter, r http.Request) {

}
