package server

import (
	"net/http"
	"os"
)

func Run() {
	assets := http.FileServer(http.Dir("./web"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	mux.HandleFunc("GET /assets/", assets.ServeHTTP)

	http.ListenAndServe(os.Getenv("APP_LISTEN_ADDR"), mux)
}
