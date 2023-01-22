package repository

import "moonlay_technical_test/model"

type CategoryRepository interface {
	FindById(id string) (*model.Category, error)
	// Find() (*[]model.Category, error)
}
