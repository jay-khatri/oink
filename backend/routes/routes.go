package routes

import (
	"encoding/json"
	"net/http"

	"backend/crypto"
	"backend/database"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 1) Get the username and password (+ do validation)
// 2) store a salted version of the password in the database
// 3) Return success/failure
func SignUp(w http.ResponseWriter, r *http.Request) {
	// Get the JSON body and decode into credentials,
	// if the structure of the body is wrong, return an HTTP error
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO(jay-khatri): validate creds here.
	// Add the user to the database.
	database.DB.Create(&database.User{
		PasswordSalted: crypto.HashAndSaltPwd(creds.Password),
		Email:          creds.Email,
	})
	// TODO(jaykhatri): Send an email for verification.
	w.WriteHeader(http.StatusOK)
}

// 1) Get the username and password.
// 2) Get the password hash for the request's email, and compare it to the request's password
// 3) Return success/failure based on the comparison.
func SignIn(w http.ResponseWriter, r *http.Request) {
	// Get the JSON body and decode into credentials,
	// if the structure of the body is wrong, return an HTTP error
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the user corresponding to the email.
	var user database.User
	database.DB.First(&user, "email = ?", creds.Email)
	// TODO(jay-khatri) review http status codes and return
	// the right ones.
	// TODO(jay-khatri) find a better way of comparison.
	if user == (database.User{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !crypto.ComparePwds(user.PasswordSalted, creds.Password) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// TODO(jaykhatri): Return the token here.
	w.WriteHeader(http.StatusOK)
}
