package main

import (
	"fmt"
	"net/http"
	"os"

	"PatientData/restAPI/driver"
	ph "PatientData/restAPI/handler/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	dbName := os.Getenv("patients")
	dbPass := os.Getenv("root1234")
	dbHost := os.Getenv("127.0.0.1")
	dbPort := os.Getenv("3306")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", postRouter(pHandler))
	})

	fmt.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)
}

func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	// r.Get("/{id:[0-9]+}", pHandler.GetByID)
	// r.Post("/", pHandler.Create)
	// r.Put("/{id:[0-9]+}", pHandler.Update)
	// r.Delete("/{id:[0-9]+}", pHandler.Delete)
	fmt.Print(r)

	return r
}
