package productcontroller

import (
	"fb-service/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{"id": 1, "nama_produk": "Product 1", "price": 100.0},
		{"id": 2, "nama_produk": "Product 2", "price": 200.0},
		{"id": 3, "nama_produk": "Product 3", "price": 300.0},
	}

	helper.ResponseJson(w, http.StatusOK, data)
}
