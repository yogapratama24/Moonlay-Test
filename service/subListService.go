package service

import (
	"moonlay_technical_test/model"
	"moonlay_technical_test/repository"
	"time"
)

type SubListService interface {
	GetSubList(listId int) (subList []model.SubList, err error)
	GetDetailSubList(Id int) (subList model.SubList, err error)
	CreateSubList(subList model.SubListCreateRequest) error
	UpdateSubList(subList model.SubListUpdateRequest) error
	DeleteSubList(Id int) error
}

type subListService struct {
	subListRepository repository.SubListRepository
}

func NewSubListService(repository repository.SubListRepository) *subListService {
	return &subListService{repository}
}

func (s *subListService) GetSubList(listId int) (subList []model.SubList, err error) {
	subList, err = s.subListRepository.GetSubList(listId)
	if err != nil {
		return
	}

	return
}

func (s *subListService) GetDetailSubList(Id int) (subList model.SubList, err error) {
	subList, err = s.subListRepository.GetDetailSubList(Id)
	if err != nil {
		return
	}

	return
}

func (s *subListService) CreateSubList(request model.SubListCreateRequest) error {
	subListCreate := model.SubList{
		Title:       request.Title,
		Description: request.Description,
		File:        request.FilePath,
		ListId:      request.List_id,
	}
	if err := s.subListRepository.CreateSubList(subListCreate); err != nil {
		return err
	}

	return nil
}

func (s *subListService) UpdateSubList(request model.SubListUpdateRequest) error {
	subListUpdate := model.SubList{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		File:        request.FilePath,
		ListId:      request.List_id,
		UpdatedAt:   time.Now(),
	}

	if err := s.subListRepository.UpdateSubList(subListUpdate); err != nil {
		return err
	}

	return nil
}

func (s *subListService) DeleteSubList(Id int) error {
	if err := s.subListRepository.DeleteSubList(Id); err != nil {
		return err
	}

	return nil
}
