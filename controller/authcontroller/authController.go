package authcontroller

import (
	"encoding/json"
	"fb-service/helper"
	"fb-service/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

	//input request user
	var inputUser models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputUser); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
	}
	defer r.Body.Close()

	//hash password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), bcrypt.DefaultCost)
	inputUser.Password = string(hashPassword)

	//insert to database
	if err := models.DB.Create(&inputUser).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
	}

	response := map[string]string{"message": "Create data successfully"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
