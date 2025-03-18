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

func (pu *ProductUseCase) EditProduct(fields []string, values []interface{},id string) (error) {

	err := pu.repository.EditProduct(fields,values,id)
	if err != nil {
		panic(err)
	}
	return err
}

func (pu *ProductUseCase) RemoveProduct(id int) (error) {
	err := pu.repository.RemoveProduct(id);
	if err != nil {
		panic(err)
	}
	return err 
}


