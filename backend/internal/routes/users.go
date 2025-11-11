package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/db"
	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/es"
)

func (h *Handler) CreateUserES(w http.ResponseWriter, req *http.Request) {
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
		"message": "user created",
	})

}

func (h *Handler) CreateUserPG(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var r db.User
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, "invalid req body", http.StatusBadRequest)
		return
	}

	db.InsertUser(r, req.Context(), h.PGpool)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "user created",
	})

}

func (h *Handler) SearchUsersByLocation(w http.ResponseWriter, req *http.Request) {

	// post because we are sending some amount of data
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var r es.SearchUserRequest
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, "Invalid Input format", http.StatusBadRequest)
		return
	}

	res, err := es.SearchUsersByLocation(req.Context(), h.ESConn, r)
	if err != nil {
		http.Error(w, "Couldnt search users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
