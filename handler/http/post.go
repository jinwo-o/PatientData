package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jinwo-o/PatientData/driver"
	repository "github.com/jinwo-o/PatientData/repository"
	post "github.com/jinwo-o/PatientData/repository/post"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

// Trying to enable CORS
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Fetch all post data
func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)
	enableCors(&w)
	respondwithJSON(w, http.StatusOK, payload)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
