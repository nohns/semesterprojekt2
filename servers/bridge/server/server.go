package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//Define a server that can be used to communicate with the rust service on the raspberry pi over http


func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/v1/certificate", certificateHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}

}

func certificateHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: implement

	//The rust client will send a request to this endpoint with the following body:
	// {
	// 	"certificate": "base64 encoded certificate",
	// 	"private_key": "base64 encoded private key"
	// }

}

// Should probaly revoke the certificate when the user logs out