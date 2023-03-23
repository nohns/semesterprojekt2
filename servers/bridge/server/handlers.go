package server

import "net/http"

func certificateHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: implement

	//The rust client will send a request to this endpoint with the following body:
	// {
	// 	"certificate": "base64 encoded certificate",
	// 	"private_key": "base64 encoded private key"
	// }

}