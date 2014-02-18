package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kir-dev/torpedo/engine"
	"github.com/kir-dev/torpedo/util"
	"html/template"
	"net/http"
	"strconv"
)

const (
	WELCOME_TEMPLATE     = "welcome"
	VIEW_TEMPLATE        = "view"
	ERROR_TEMPLATE       = "error"
	SHOOT_TEMPLATE       = "shoot"
	HISTORYLIST_TEMPLATE = "history"
	PLAYER_ID_COOKIE     = "id"
	GAME_ID_COOKIE       = "gid"
)

var (
	// registered routes, with handlers
	getRoutesMap = map[string]http.HandlerFunc{
		"/":            rootHandler,
		"/view":        viewHandler,
		"/shoot":       shootHandler,
		"/quit":        quitHandler,
		"/games":       historyHandler,
		"/games/{gid}": historyHandler,
	}
	postRoutesMap = map[string]http.HandlerFunc{
		"/join":  joinHandler,
		"/shoot": shootHandler,
	}

	// template cache
	templates = template.New("root")
)

type ShootResult struct {
	Feedback engine.HitResult
	Game     *engine.Game
	Player   *engine.Player
}

// register handlers for routes
func init() {

	r := mux.NewRouter()
	for path, handler := range getRoutesMap {
		r.HandleFunc(path, handler).Methods("GET")
	}
	for path, handler := range postRoutesMap {
		r.HandleFunc(path, handler).Methods("POST")
	}
	http.Handle("/", r)

	templates.Funcs(utilFuncMap())
	templates = template.Must(templates.ParseGlob("template/*.html.go"))
}

/******* handlers ******/

// welcome page with registration
// also every not registered routes will throw a 404
func rootHandler(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: GAME_ID_COOKIE, Value: currentGame.Id, HttpOnly: true})
	renderTemplate(w, WELCOME_TEMPLATE, nil)
}

// join handler
func joinHandler(rw http.ResponseWriter, req *http.Request) {
	if !canJoin(req) {
		renderError(rw, util.Errorf("You are already playing! You cannot join twice."))
		return
	}

	player := engine.NewPlayer(req.FormValue("username"))
	player.IsBot = covertCheckboxValueToBool(req.FormValue("is_robot"))

	err := player.Join(currentGame)
	if err != nil {
		util.LogWarn("Player could not join. Cause: %s", err.Error())
		renderTemplate(rw, ERROR_TEMPLATE, errorView{
			// TODO: move to properties file
			"Játékos nem tudott csatlakozni.",
			err.Error(),
			util.IsDev(),
		})
	} else {
		http.SetCookie(rw, &http.Cookie{Name: PLAYER_ID_COOKIE, Value: player.Id, HttpOnly: true})
		http.Redirect(rw, req, "/shoot", http.StatusFound)
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

	var shootResult ShootResult
	shootResult.Game = currentGame
	shootResult.Player = currentGame.GetPlayerById(cookie.Value)

	if r.Method == "POST" {
		// check if the player is up
		if currentGame.CurrentPlayerId != cookie.Value {
			// -1 : it's not your turn
			fmt.Fprint(w, "-1")
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

		shootResult.Feedback = currentGame.Shoot(row, col)
		util.LogInfo("Player shot at (%s) with result: %s", engine.RowColToS(row, col), shootResult.Feedback)
		fmt.Fprint(w, shootResult.Feedback)
		return
	}

	renderTemplate(w, SHOOT_TEMPLATE, shootResult)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	urlPattern := mux.Vars(r)
	gid := urlPattern["gid"]
	if gid != "" {
		var selectedGame *engine.Game = nil
		for _, game := range History {
			if game.Id == gid {
				selectedGame = game
				break
			}
		}
		renderTemplate(w, VIEW_TEMPLATE, selectedGame)
	} else {
		renderTemplate(w, HISTORYLIST_TEMPLATE, History)
	}
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
			util.IsDev(),
		})
		return true
	}
	return false

}

func renderError(w http.ResponseWriter, err error) {
	renderTemplate(w, ERROR_TEMPLATE, errorView{
		err.Error(),
		err.Error(),
		util.IsDev(),
	})
}

// returns the template name that can be used for rendering
func getFullTempalteName(tmplName string) string {
	return tmplName + ".html.go"
}

// Writes the tempalte to the response. If we are running in a development mode,
// it re-reads the template from the file every time. Otherwise it uses the
// cached version.
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	if util.IsDev() {
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

type errorView struct {
	Title   string
	Message string
	IsDev   bool
}

// A user can join if he has the correct game id and either he does not have
// a player id or his player id is not in the game.
func canJoin(req *http.Request) bool {
	gid, errGid := req.Cookie(GAME_ID_COOKIE)
	pid, errPid := req.Cookie(PLAYER_ID_COOKIE)

	// has game id and its
	if errGid == nil && gid.Value == currentGame.Id {
		// has player id in cookie and current game has the player
		if errPid == nil && currentGame.GetPlayerById(pid.Value) != nil {
			return false
		}

		return true
	}

	return false
}
