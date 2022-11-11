package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"infoblog/internal/models"
	"net/http"
	"strconv"
)

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
func (app *application) post(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "samplePost.html", data)
}

func (app *application) blogView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	infoBlog, err := app.infoBlogs.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.InfoBlog = infoBlog
	app.render(w, http.StatusOK, "samplePost.html", data)
}

func (app *application) blogPost(w http.ResponseWriter, r *http.Request) {

}
