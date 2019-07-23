package handler

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinwo-o/PatientData/driver"
	models "github.com/jinwo-o/PatientData/models"
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

// Fetch all patient data
func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)
	enableCors(&w)
	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID will fetch the patient by ID
func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		fmt.Println("Error: ID does not exist")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new Patient
func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.Create(r.Context(), &post)
	// ?
	fmt.Println(newID)
	if err != nil {
		fmt.Println("Error: Unable to create patient")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a patient by id
func (p *Post) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		fmt.Println("Error: Unable to update patient")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a patient
func (p *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		fmt.Println("Error: Unable to delete patient")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
