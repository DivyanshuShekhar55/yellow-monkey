package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/es"
)

func (h *Handler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	res, err := es.GetAllGroups(r.Context(), h.ESConn)

	if err != nil {
		http.Error(w, "error getting all groups", http.StatusInternalServerError)
		log.Printf("error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) CreateUserIndex() {}
