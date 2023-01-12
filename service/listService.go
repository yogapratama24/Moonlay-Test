package service

import (
	"moonlay_technical_test/helper"
	"moonlay_technical_test/model"
	"moonlay_technical_test/repository"
	"time"
)

type ListService interface {
	GetList(request helper.Pagination) (pagination helper.Pagination, err error)
	// GetListRelational() (lists []model.ListResponseRelational, err error)
	GetListRelational() (lists []model.List, err error)
	GetDetailList(Id int) (list model.List, err error)
	CreateList(list model.ListCreateRequest) error
	UpdateList(list model.ListUpdateRequest) error
	DeleteList(Id int) error
}

type listService struct {
	listRepository repository.ListRepository
}

func NewListService(repository repository.ListRepository) *listService {
	return &listService{repository}
}

func (s *listService) GetList(request helper.Pagination) (pagination helper.Pagination, err error) {
	pagination, err = s.listRepository.GetList(request)
	if err != nil {
		return
	}

	return
}

func (s *listService) GetListRelational() (data []model.List, err error) {
	data, err = s.listRepository.GetListRelational()
	if err != nil {
		return
	}

	return
}

func (s *listService) GetDetailList(Id int) (list model.List, err error) {
	list, err = s.listRepository.GetDetailList(Id)
	if err != nil {
		return
	}

	return
}

func (s *listService) CreateList(request model.ListCreateRequest) error {
	listCreate := model.List{
		Title:       request.Title,
		Description: request.Description,
		File:        request.FilePath,
	}
	if err := s.listRepository.CreateList(listCreate); err != nil {
		return err
	}

	return nil
}

func (s *listService) UpdateList(request model.ListUpdateRequest) error {
	listUpdate := model.List{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		File:        request.FilePath,
		UpdatedAt:   time.Now(),
	}

	if err := s.listRepository.UpdateList(listUpdate); err != nil {
		return err
	}

	return nil
}

func (s *listService) DeleteList(Id int) error {
	if err := s.listRepository.DeleteList(Id); err != nil {
		return err
	}

	return nil
}
