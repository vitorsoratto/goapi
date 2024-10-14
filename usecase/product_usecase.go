package usecase

import (
	"goapi/model"
	"goapi/repository"
)

type ProductUsecase struct {
	productRepository *repository.ProductRepository
}

func NewProductUsecase(repository *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepository: repository,
	}
}

func (u *ProductUsecase) GetProducts() ([]model.Product, error) {
	products, err := u.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *ProductUsecase) GetProductByID(productID int) (model.Product, error) {
	product, err := u.productRepository.GetProductByID(productID)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (u *ProductUsecase) InsertProduct(product model.Product) (model.Product, error) {
	productID, err := u.productRepository.InsertProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil
}
