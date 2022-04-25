package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Faiber-Barragan/twittor/db"
	"github.com/Faiber-Barragan/twittor/models"
)

//SignUp is the function to create in the DB the register of user
func SignUp(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in received data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "User Email is required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	_, found, _ := db.CheckUserAlreadyExists(t.Email)
	if found {
		http.Error(w, "Already exists user sign up with that email", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to perform a user registration: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "Failed to insert user record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
