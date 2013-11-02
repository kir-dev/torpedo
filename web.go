package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	WELCOME_TEMPLATE = "welcome"
	VIEW_TEMPLATE    = "view"
	ERROR_TEMPLATE   = "error"
)

var (
	// registered routes, with handlers
	routesMap = map[string]func(http.ResponseWriter, *http.Request){
		"/":     rootHandler,
		"/join": joinHandler,
		"/view": viewHandler,
	}
	// template cache
	templates = template.New("root")
	// to avoid initialization loop compilation error...
	routes = make([]string, 0, 5)
)

// register handlers for routes
func init() {
	for route, handler := range routesMap {
		http.HandleFunc(route, handler)
		routes = append(routes, route)
	}

	templates.Funcs(utilFuncMap())
	templates = template.Must(templates.ParseGlob("template/*.html.go"))
}

// welcome page with registration
// also every not registered routes will throw a 404
func rootHandler(w http.ResponseWriter, req *http.Request) {
	err := check404(w, req)
	if err != nil {
		log.Print(err)
		return
	}

	renderTemplate(w, WELCOME_TEMPLATE, nil)
}

// join handler
func joinHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		player := Player{
			req.FormValue("username"),
			covertCheckboxValueToBool(req.FormValue("is_robot")),
			nil,
		}
		err := join(&player)
		if err != nil {
			logWarn("Player could not join. Cause: %s", err.Error())
			renderTemplate(rw, ERROR_TEMPLATE, errorView{
				// TODO: move to properties file
				"Játékos nem tudott csatlakozni.",
				err.Error(),
				isDev(),
			})
		} else {
			http.Redirect(rw, req, "/view", http.StatusFound)
		}

	} else {
		http.NotFound(rw, req)
		log.Print(messageFor404(req))
	}
}

// view handler -- shows information about the current game
func viewHandler(rw http.ResponseWriter, req *http.Request) {
	renderTemplate(rw, VIEW_TEMPLATE, currentGame)
}

func check404(w http.ResponseWriter, req *http.Request) error {
	found := false
	for _, route := range routes {
		if req.RequestURI == route {
			found = true
			break
		}
	}

	if !found {
		http.NotFound(w, req)
		return errors.New(messageFor404(req))
	}

	return nil
}

// returns the template name that can be used for rendering
func getFullTempalteName(tmplName string) string {
	return tmplName + ".html.go"
}

// Writes the tempalte to the response. If we are running in a development mode,
// it re-reads the template from the file every time. Otherwise it uses the
// cached version.
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	if isDev() {
		tmpl := template.New(getFullTempalteName(tmplName))
		tmpl.Funcs(utilFuncMap())
		//tmpl, err := template.ParseFiles("template/" + getFullTempalteName(tmplName))
		tmpl, err := tmpl.ParseFiles("template/" + getFullTempalteName(tmplName))
		if err != nil {
			fmt.Fprintf(w, "Error in template:\n%s", err.Error())

		} else {
			tmpl.Execute(w, data)
		}
		return

	}
	templates.ExecuteTemplate(w, getFullTempalteName(tmplName), data)
}

func covertCheckboxValueToBool(value string) bool {
	if value == "on" {
		return true
	}
	return false
}

// creates the message for 404
func messageFor404(req *http.Request) string {
	return fmt.Sprintf("(404) Not found: [%s] %s", req.Method, req.RequestURI)
}
