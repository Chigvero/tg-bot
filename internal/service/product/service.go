package product

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProduct
}

func (s *Service) Get(idx int) (Product, error) {
	if idx < 0 || idx > len(allProduct) {
		return Product{}, errors.New("need correct Id")
	}
	return allProduct[idx], nil
}
