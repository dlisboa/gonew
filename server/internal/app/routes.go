package app

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.home)

	return mux
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
