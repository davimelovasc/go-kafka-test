package usecase

import "github.com/davimelovasc/go-simple-api/internal/entity"

type ProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUseCase struct {
	entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ProductOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var productsOutput []*ProductOutputDto
	for _, product := range products {
		productsOutput = append(productsOutput, &ProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}
