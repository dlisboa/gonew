package server

import (
	"net/http"
	"os"
)

func Run() error {
	assets := http.FileServer(http.Dir("./web"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	mux.Handle("GET /static/", http.StripPrefix("/static/", assets))

	return http.ListenAndServe(os.Getenv("APP_LISTEN_ADDR"), mux)
}
