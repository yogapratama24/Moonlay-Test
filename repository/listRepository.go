package repository

import (
	"log"
	"moonlay_technical_test/helper"
	"moonlay_technical_test/model"

	"gorm.io/gorm"
)

type ListRepository interface {
	GetList(request helper.Pagination) (pagination helper.Pagination, err error)
	GetListRelational() (lists []model.List, err error)
	GetDetailList(Id int) (list model.List, err error)
	CreateList(list model.List) error
	UpdateList(list model.List) error
	DeleteList(Id int) error
}

type listRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *listRepository {
	return &listRepository{db}
}

func (r *listRepository) GetList(request helper.Pagination) (pagination helper.Pagination, err error) {
	db := r.db
	var (
		lists      = []model.List{}
		queryTitle string
	)
	if request.Title != "" {
		db.Scopes(helper.Paginate(lists, &request, db)).Where("title LIKE ?", "%"+request.Title+"%").Find(&lists)
		pagination.Rows = lists
	} else {
		db.Scopes(helper.Paginate(lists, &request, db)).Where(queryTitle).Find(&lists)
		pagination.Rows = lists
	}
	// if err = db.Find(&lists).Error; err != nil {
	// 	log.Printf("Error get data list err: %s", err)
	// 	return
	// }
	return
}

func (r *listRepository) GetListRelational() (lists []model.List, err error) {
	db := r.db
	if err = db.Preload("SubList").Find(&lists).Error; err != nil {
		log.Printf("Error get data list relational with err: %s", err)
		return
	}
	return
}

func (r *listRepository) GetDetailList(Id int) (list model.List, err error) {
	db := r.db
	if err = db.Find(&list, Id).Error; err != nil {
		log.Printf("Error get data detail list with err: %s", err)
		return
	}
	return
}

func (r *listRepository) CreateList(list model.List) error {
	db := r.db
	err := db.Create(&list)
	if err.Error != nil {
		log.Printf("Error create data list with err: %v", err)
		return err.Error
	}
	return nil
}

func (r *listRepository) UpdateList(list model.List) error {
	db := r.db
	err := db.Model(list).Updates(list)
	if err.Error != nil {
		log.Printf("Error update data list with err: %v", err)
		return err.Error
	}

	return nil
}

func (r *listRepository) DeleteList(Id int) error {
	db := r.db
	err := db.Delete(model.List{}, Id)
	if err.Error != nil {
		log.Printf("Error delete data list with err: %v", err)
		return err.Error
	}

	return nil
}
