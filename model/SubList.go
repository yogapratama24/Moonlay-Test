package model

import (
	"time"
)

type SubList struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:100"`
	Description string    `json:"description"`
	File        string    `json:"file" gorm:"size:300"`
	ListId      int       `json:"list_id"`
	List        *List     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"list"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type SubListRelational struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:100"`
	Description string    `json:"description"`
	File        string    `json:"file" gorm:"size:300"`
	ListId      int       `json:"list_id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type SubListRequest struct {
	Id          int    `json:"id" query:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
	List        string `json:"list"`
	ListId      int    `json:"list_id" query:"list_id" validate:"required"`
}

type SubListCreateRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title" query:"title" validate:"required"`
	Description string `json:"description" query:"description" validate:"required"`
	FilePath    string `json:"file_path"`
	List_id     int    `json:"list_id" query:"list_id" validate:"required"`
}

type SubListUpdateRequest struct {
	Id          int    `json:"id" query:"id" validate:"required"`
	Title       string `json:"title" query:"title" validate:"required"`
	Description string `json:"description" query:"description" validate:"required"`
	FilePath    string `json:"file" query:"file" gorm:"size:300"`
	List_id     int    `json:"list_id" query:"list_id"`
}

type SubListDeleteRequest struct {
	Id int `json:"id" param:"id" query:"id" validate:"required"`
}
