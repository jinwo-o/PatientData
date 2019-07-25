package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	// good separation into packages!
	"github.com/jinwo-o/PatientData/driver"
	ph "github.com/jinwo-o/PatientData/handler/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// add command line arguments so you can chose the specific port to bind to
	// (with the default value being 8080), IP address or hostname for the
	// MySQL database you wish to use

	port := flag.String("port", "8080", "port bound to listening service")
	ipAddress := flag.String("ip", "0.0.0.0", "IP address to bind service to")
	flag.Parse()

	connection, err := driver.ConnectSQL()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/patients", postRouter(pHandler))
	})

	// here you can use the arguments passed at the command line instead of
	// hard-coded values
	listeningMessage := "Server listening on port: " + *port
	fmt.Println(listeningMessage)

	ipAndPortService := *ipAddress + ":" + *port
	//http.ListenAndServe("0.0.0.0:8080", r)
	http.ListenAndServe(ipAndPortService, r)
}

// Notes 2017-07-25
// when dealing with endpoints, if you're looking up a single patient the
// endpoint should really be: /patient/5 and not /patients/5
// just a small standard convention thing
func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)

	return r
}
