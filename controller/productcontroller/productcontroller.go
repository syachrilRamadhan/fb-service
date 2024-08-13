package productcontroller

import (
	"encoding/json"
	"fb-service/helper"
	"fb-service/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	var products []models.Product

	models.DB.Find(&products)
	data := map[string]interface{}{
		"data": products,
	}

	helper.ResponseJson(w, http.StatusOK, data)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	params := mux.Vars(r)

	models.DB.Where("id = ?", params["id"]).First(&product)

	if product.Id == 0 {
		response := map[string]string{"message": "Data not found"}
		helper.ResponseJson(w, http.StatusNotFound, response)
		return
	}

	data := map[string]interface{}{
		"data": product,
	}

	helper.ResponseJson(w, http.StatusOK, data)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	// Create product into database
	result := models.DB.Create(&product)
	if result.Error != nil {
		response := map[string]string{"message": result.Error.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"message": "Create product successfully",
		"data":    product,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	params := mux.Vars(r)
	models.DB.Where("id = ?", params["id"]).First(&product)

	if product.Id == 0 {
		response := map[string]string{"message": "Data not found"}
		helper.ResponseJson(w, http.StatusNotFound, response)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	models.DB.Save(&product)

	response := map[string]interface{}{
		"message": "Update product successfully",
		"data":    product,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	models.DB.Where("id = ?", params["id"]).First(&product)

	if product.Id == 0 {
		response := map[string]string{"message": "Data not found"}
		helper.ResponseJson(w, http.StatusNotFound, response)
		return
	}

	models.DB.Delete(&product)

	response := map[string]string{"message": "Delete product successfully"}
	helper.ResponseJson(w, http.StatusOK, response)
}
