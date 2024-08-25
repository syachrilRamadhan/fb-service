package authcontroller

import (
	"encoding/json"
	"fb-service/config"
	"fb-service/helper"
	"fb-service/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var inputUser models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputUser); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//get data user
	var user models.User

	if err := models.DB.Select("username, password").Where("username = ?", inputUser.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username tidak terdaftar"}
			helper.ResponseJson(w, http.StatusNotFound, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	//matching passwords match
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah!"}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	//create jwt token
	expTime := time.Now().Add(time.Minute * 60)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "fb-service",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//declare signing algorithm
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	//set token into cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]interface{}{
		"message": "Login Success",
		"token":   token,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {

	//input request user
	var inputUser models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputUser); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//hash password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), bcrypt.DefaultCost)
	inputUser.Password = string(hashPassword)

	//insert to database
	if err := models.DB.Create(&inputUser).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Create data successfully"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	//delete token
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout Success"}
	helper.ResponseJson(w, http.StatusOK, response)
}
