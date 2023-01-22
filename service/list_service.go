package service

import (
	"errors"
	"moonlay_technical_test/helper"
	"moonlay_technical_test/model"
	"moonlay_technical_test/repository"
)

type ListService interface {
	GetList(request helper.Pagination) (pagination helper.Pagination, err error)
	// GetListRelational() (lists []model.ListResponseRelational, err error)
	GetListRelational() (lists []model.List, err error)
	GetDetailList(Id int) (*model.List, error)
	CreateList(list model.ListCreateRequest) error
	UpdateList(list model.ListUpdateRequest) error
	DeleteList(Id int) error
}

type ListServices struct {
	ListRepository repository.ListRepository
}

func NewListService(repository repository.ListRepository) *ListServices {
	return &ListServices{repository}
}

func (s *ListServices) GetList(request helper.Pagination) (pagination helper.Pagination, err error) {
	pagination, err = s.ListRepository.GetList(request)
	if err != nil {
		return
	}

	return
}

func (s *ListServices) GetListRelational() (data []model.List, err error) {
	data, err = s.ListRepository.GetListRelational()
	if err != nil {
		return
	}

	return
}

func (s *ListServices) GetDetailList(Id int) (*model.List, error) {
	list, err := s.ListRepository.GetDetailList(Id)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return list, errors.New("error get list")
	} else {
		return list, nil
	}

}

func (s *ListServices) CreateList(request model.ListCreateRequest) error {
	listCreate := model.List{
		Id:          1,
		Title:       request.Title,
		Description: request.Description,
		File:        request.FilePath,
	}
	if err := s.ListRepository.CreateList(listCreate); err != nil {
		return err
	}

	return nil
}

func (s *ListServices) UpdateList(request model.ListUpdateRequest) error {
	listUpdate := model.List{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		File:        request.FilePath,
		// UpdatedAt:   time.Now(),
	}

	if err := s.ListRepository.UpdateList(listUpdate); err != nil {
		return err
	}

	return nil
}

func (s *ListServices) DeleteList(Id int) error {
	if err := s.ListRepository.DeleteList(Id); err != nil {
		return err
	}

	return nil
}
