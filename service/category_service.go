package service

import (
	"errors"
	"moonlay_technical_test/model"
	"moonlay_technical_test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*model.Category, error) {
	category, _ := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("error get category")
	} else {
		return category, nil
	}
}
