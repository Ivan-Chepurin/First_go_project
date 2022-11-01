package main

import (
	"errors"
	"fmt"
	"main/pkg/models"
	"net/http"
	"strconv"
)

var (
	unknownErr = errors.New("внутренняя ошибка сервера")
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.Snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n\n", snippet)
	}

	//files := []string{
	//	"./ui/html/home.page.tmpl.html",
	//	"./ui/html/base.layout.tmpl.html",
	//	"./ui/html/footer.partial.tmpl.html",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//err = ts.Execute(w, nil)
	//if err != nil {
	//	app.serverError(w, err)
	//}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.Snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
		return

	}
	_, err = fmt.Fprintf(w, "%v", snippet)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "История про улитку"
	content := "Улитка выползла из раковины,\nвытянула рожки,\nи опять подобрала их."
	expires := "7"

	id, err := app.Snippets.Insert(title, content, expires)
	if err != nil {
		w.Write([]byte(unknownErr.Error()))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
