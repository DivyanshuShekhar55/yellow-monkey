package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/es"
)

func (h *Handler) CreateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var r es.User
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, "invalid req body", http.StatusBadRequest)
		return
	}

	es.PutUser(r, h.ESConn)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "group created",
	})

}

func(h *Handler) SearchUsersByLocation(){
	
}