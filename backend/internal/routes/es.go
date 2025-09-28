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

func (h *Handler) CreateGroup(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "method not alllowed", http.StatusMethodNotAllowed)
		return
	}

	var r es.Group
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	es.InsertGroup(req.Context(), h.ESConn, &r, "wait_for")
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "group created",
	})
}
