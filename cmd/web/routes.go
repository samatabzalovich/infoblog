package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/infoBlog/about", dynamic.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/users/login", dynamic.ThenFunc(app.login))
	router.Handler(http.MethodPost, "/users/login", dynamic.ThenFunc(app.userLoginPost))
	router.Handler(http.MethodGet, "/users/signup", dynamic.ThenFunc(app.signup))
	router.Handler(http.MethodPost, "/users/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodGet, "/infoBlog/contact", dynamic.ThenFunc(app.contact))
	router.Handler(http.MethodGet, "/infoBlog/samplePost", dynamic.ThenFunc(app.post))
	router.Handler(http.MethodGet, "/infoblogs/view/:id", dynamic.ThenFunc(app.blogView))
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodPost, "/users/logout", protected.ThenFunc(app.userLogoutPost))
	//router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	//router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
