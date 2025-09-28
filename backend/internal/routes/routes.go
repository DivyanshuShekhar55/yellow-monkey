package routes

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/v9"
)

type Handler struct {
	ESConn *elasticsearch.Client
}

func NewHandler(esconn *elasticsearch.Client) *Handler {
	return &Handler{
		ESConn: esconn,
	}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func( w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/groups", h.GetAllGroups)
}