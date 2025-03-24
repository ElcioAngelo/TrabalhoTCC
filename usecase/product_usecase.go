package usecase

import (
	"github.comElcioAngelo/TrabalhoTCC.git/model"
	"github.comElcioAngelo/TrabalhoTCC.git/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (error) {
	err := pu.repository.CreateProduct(product)
	if (err != nil) {
		panic(err)
	}

	return err
}

func (pu *ProductUseCase) GetProductById(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func(pu *ProductUseCase) EditProductName(id_product int, value_to_update string) (error) {
	err := pu.repository.EditProductName(id_product,value_to_update)
	if err != nil {
		return err
	}

	return err
}

func(pu *ProductUseCase) EditProductCategory(id_product int, value_to_update string) (error)  {
	err := pu.repository.EditProductCategory(id_product,value_to_update)
	if err != nil {
		return err
	}

	return err
}

func(pu *ProductUseCase) EditProductPrice(id_product int, value_to_update string) (error)  {
	err := pu.repository.EditProductPrice(id_product,value_to_update)
	if err != nil {
		return err
	}

	return err
}


