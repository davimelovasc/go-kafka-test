package web

import (
	"encoding/json"
	"net/http"

	"github.com/davimelovasc/go-simple-api/internal/usecase"
)

type ProductHandler struct {
	*usecase.CreateProductUseCase
	*usecase.ListProductsUseCase
}

func NewProductHandler(createProductUseCase *usecase.CreateProductUseCase, listProductUseCase *usecase.ListProductsUseCase) *ProductHandler {
	return &ProductHandler{
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductUseCase,
	}
}

func (p *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
