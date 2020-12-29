package handlers

import (
	"fmt"
	"net/http"

	"github.com/el10savio/woot-crdt/woot"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	// WString is the WString
	// data structure initialized
	WString woot.WString
)

func init() {
	WString = woot.Initialize()
}

// Route defines the Mux
// router individual route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a collection
// of individual Routes
var Routes = []Route{
	{"/woot", "GET", WOOT},
	{"/woot/list", "GET", List},
	{"/woot/add", "POST", Add},
	{"/woot/delete", "POST", Delete},
	{"/woot/sync/add", "POST", SyncAdd},
	{"/woot/sync/delete", "POST", SyncDelete},
}

// WOOT is the handler for the path "/woot"
func WOOT(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World WOOT Node\n")
}

// Logger is the middleware to
// log the incoming request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")
		next.ServeHTTP(w, r)
	})
}

// Router returns a mux router
func Router() *mux.Router {
	// Initialize Router
	router := mux.NewRouter()

	// Instantiate Routes
	for _, route := range Routes {
		router.HandleFunc(
			route.Path,
			route.Handler,
		).Methods(route.Method)
	}

	// Add static file serve handler
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./handlers/public")))

	// Enable Router Logger
	router.Use(Logger)

	return router
}
