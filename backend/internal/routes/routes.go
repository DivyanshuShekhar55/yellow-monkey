package routes

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	ESConn *elasticsearch.Client
	PGpool *pgxpool.Pool
}

func NewHandler(esconn *elasticsearch.Client, pgpool *pgxpool.Pool) *Handler {
	return &Handler{
		ESConn: esconn,
		PGpool : pgpool,
	}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func( w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/groups", h.GetAllGroups)
	mux.HandleFunc("POST /group", h.CreateGroup)
	mux.HandleFunc("POST /groups/search", h.SearchGroups)
	mux.HandleFunc("POST /user/search-by-loc", h.SearchUsersByLocation)
	mux.HandleFunc("POST /user", h.CreateUserES)
}