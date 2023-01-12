package repository

import (
	"log"
	"moonlay_technical_test/model"

	"gorm.io/gorm"
)

type SubListRepository interface {
	GetSubList(listId int) (subList []model.SubList, err error)
	GetDetailSubList(Id int) (subList model.SubList, err error)
	CreateSubList(subList model.SubList) error
	UpdateSubList(subList model.SubList) error
	DeleteSubList(Id int) error
}

type subListRepository struct {
	db *gorm.DB
}

func NewSubListRepository(db *gorm.DB) *subListRepository {
	return &subListRepository{db}
}

func (r *subListRepository) GetSubList(listId int) (subList []model.SubList, err error) {
	db := r.db
	if err = db.Where("list_id = ?", listId).Find(&subList).Error; err != nil {
		log.Printf("Error get data sub list err: %s", err)
		return
	}
	return
}

func (r *subListRepository) GetDetailSubList(Id int) (subList model.SubList, err error) {
	db := r.db
	if err = db.Find(&subList, Id).Error; err != nil {
		log.Printf("Error get data detail sub list with err: %s", err)
		return
	}
	return
}

func (r *subListRepository) CreateSubList(subList model.SubList) error {
	db := r.db
	err := db.Create(&subList)
	if err.Error != nil {
		log.Printf("Error create data sub list with err: %v", err)
		return err.Error
	}
	return nil
}

func (r *subListRepository) UpdateSubList(subList model.SubList) error {
	db := r.db
	err := db.Model(subList).Updates(subList)
	if err.Error != nil {
		log.Printf("Error update data sub list with err: %v", err)
		return err.Error
	}

	return nil
}

func (r *subListRepository) DeleteSubList(Id int) error {
	db := r.db
	err := db.Delete(model.SubList{}, Id)
	if err.Error != nil {
		log.Printf("Error delete data sub list with err: %v", err)
		return err.Error
	}

	return nil
}
