package handlers

import (
	"encoding/json"
	"net/http"
	"webhook-project/models"
	"webhook-project/services"
	"webhook-project/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	services.CreateUser(user)
	w.Write([]byte("User created"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	user, err := services.FindUser(input.Username)

	if err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	token, _ := utils.GenerateToken(user.Username)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}