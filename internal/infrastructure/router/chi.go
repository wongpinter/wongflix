package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

var (
	chiDispatcher = chi.NewRouter()
)

// var tokenAuth *jwtauth.JWTAuth

// TODO FIND OUT WHY JWT NOT WORKING!!
// func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
// 	chiDispatcher.Group(func(r chi.Router) {
// 		r.Use(jwtauth.Verifier(tokenAuth))

// 		r.Use(jwtauth.Authenticator)

// 		r.Get(uri, f)
// 	})
// }

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (c *chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server listening on port %v\n", port)
	http.ListenAndServe(port, chiDispatcher)
}
