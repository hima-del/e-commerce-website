package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
)

func startApp() {
	mapUrls()
	http.ListenAndServe(":80", r)
}
