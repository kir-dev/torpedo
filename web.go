package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	WELCOME_TEMPLATE = "welcome"
	VIEW_TEMPLATE    = "view"
	ERROR_TEMPLATE   = "error"
	SHOOT_TEMPLATE   = "shoot"
	PLAYER_ID_COOKIE = "id"
	GAME_ID_COOKIE   = "gid"
)

var (
	// registered routes, with handlers
	routesMap = map[string]http.HandlerFunc{
		"/":      rootHandler,
		"/join":  joinHandler,
		"/view":  viewHandler,
		"/shoot": shootHandler,
		"/quit":  quitHandler,
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

/******* handlers ******/

// welcome page with registration
// also every not registered routes will throw a 404
func rootHandler(w http.ResponseWriter, req *http.Request) {
	err := check404(w, req)
	if err != nil {
		log.Print(err)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: GAME_ID_COOKIE, Value: currentGame.Id, HttpOnly: true})
	renderTemplate(w, WELCOME_TEMPLATE, nil)
}

// join handler
func joinHandler(rw http.ResponseWriter, req *http.Request) {
	var canJoin bool

	gid, errGid := req.Cookie(GAME_ID_COOKIE)
	_, errPid := req.Cookie(PLAYER_ID_COOKIE)

	switch {
	case errGid != nil:
		// no game id cookie -> did not participate in any game before
		canJoin = true
	case gid.Value != currentGame.Id:
		// game id is not the current game's id, so it can join again
		canJoin = true
	case gid.Value == currentGame.Id && errPid == nil:
		// has the current game's id and the player id as well -> cannot join
		canJoin = false
	case errPid != nil:
		// does not have a player id -> automatically qualifies for join
		canJoin = true
	}

	if !canJoin {
		renderError(rw, errorf("You are already playing! You cannot join twice."))
		return
	}

	if req.Method == "POST" {
		player := newPlayer(req.FormValue("username"))
		player.IsBot = covertCheckboxValueToBool(req.FormValue("is_robot"))

		err := player.join(currentGame)
		if err != nil {
			logWarn("Player could not join. Cause: %s", err.Error())
			renderTemplate(rw, ERROR_TEMPLATE, errorView{
				// TODO: move to properties file
				"Játékos nem tudott csatlakozni.",
				err.Error(),
				isDev(),
			})
		} else {
			http.SetCookie(rw, &http.Cookie{Name: PLAYER_ID_COOKIE, Value: player.Id, HttpOnly: true})
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

// shoot handler
func shootHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(PLAYER_ID_COOKIE)

	// the play might not be registered
	if err != nil {
		// TODO: error handling
		fmt.Fprint(w, "Join the game first!")
		return
	}

	var feedback hitResult
	if r.Method == "POST" {
		// check if the player is up
		if currentGame.CurrentPlayerId != cookie.Value {
			// TODO: error handling
			fmt.Fprint(w, "It's not your turn!")
			return
		}

		colS := r.FormValue("col")
		rowS := r.FormValue("row")

		row, err := strconv.Atoi(rowS)
		if renderParseError(w, err) {
			return
		}
		// 0 based indexing
		row -= 1

		if len(colS) == 0 && renderParseError(w, errors.New("Column should be [A-Z]")) {
			return
		}
		col := int(colS[0] - 'A')

		logInfo("Player shot at (%s)", rcToS(row, col))
		feedback = currentGame.Board.shootAt(row, col, currentGame.endTurn)
	}

	renderTemplate(w, SHOOT_TEMPLATE, feedback)
}

// Makes the player quit the game, sets the player to be a bot for the rest of
// the game. There is no re-entry option.
func quitHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(PLAYER_ID_COOKIE)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	for _, player := range currentGame.Players {
		// if player exists, set it to be a bot
		if player.Id == cookie.Value {
			player.IsBot = true
			break
		}
	}
	// delete the cookie
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

/**** utility methods ****/

func renderParseError(w http.ResponseWriter, err error) bool {
	if err != nil {
		renderTemplate(w, ERROR_TEMPLATE, errorView{
			"Could not parse integer",
			err.Error(),
			isDev(),
		})
		return true
	}
	return false

}

func renderError(w http.ResponseWriter, err error) {
	renderTemplate(w, ERROR_TEMPLATE, errorView{
		err.Error(),
		err.Error(),
		isDev(),
	})
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
