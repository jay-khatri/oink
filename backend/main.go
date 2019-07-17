package main

import (
	"backend/routes"
	"log"
	"net/http"
	// "backend/crypto"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement.
	http.HandleFunc("/signup", routes.SignUp)
	http.HandleFunc("/signin", routes.SignIn)
	//TODO(jaykhatri) add a refresh endpoint for refreshing the token.

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
