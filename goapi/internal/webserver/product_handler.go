package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	entity "github.com/idevbn/imersao-fullstack-fullcycle/goapi/internal/entity"
	"github.com/idevbn/imersao-fullstack-fullcycle/goapi/internal/service"
)

type WebProductHandler struct {
	ProdcutService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProdcutService: productService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProdcutService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	product, err := wph.ProdcutService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProductsByCategoryId(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "icategryID is required", http.StatusBadRequest)
		return
	}

	products, err := wph.ProdcutService.GetProductsByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wph.ProdcutService.CreateProduct(
		product.Name,
		product.Description,
		product.CategoryID,
		product.ImageURL,
		product.Price,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
