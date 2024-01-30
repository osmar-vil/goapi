package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/osmar-vil/goapi/internal/entity"
	"github.com/osmar-vil/goapi/internal/service"
)

type WebProductHandler struct {
	service *service.ProductService
}

func NewWebProductHandler(service *service.ProductService) *WebProductHandler {
	return &WebProductHandler{service}
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	e := json.NewDecoder(r.Body).Decode(&product)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	_, e = wph.service.CreateProduct(
		product.Name,
		product.Description,
		product.ImageURL,
		product.CategoryID,
		product.Price,
	)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, e := wph.service.GetProducts()
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	product, e := wph.service.GetProduct(id)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	products, e := wph.service.GetProductsByCategoryID(id)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(products)
}
