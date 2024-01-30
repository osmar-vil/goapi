package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/osmar-vil/goapi/internal/entity"
	"github.com/osmar-vil/goapi/internal/service"
)

type WebCategoryHandler struct {
	categoryService *service.CategoryService
}

func NewWebCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{categoryService}
}

func (wch *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, e := wch.categoryService.GetCategories()
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (wch *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	category, e := wch.categoryService.GetCategory(id)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (wch *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	e := json.NewDecoder(r.Body).Decode(&category)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	result, e := wch.categoryService.CreateCategory(category.Name)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
