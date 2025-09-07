package routes

import "net/http"

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func( w http.ResponseWriter, r *http.Request,) {
		w.Write([]byte("ok"))
	})
}