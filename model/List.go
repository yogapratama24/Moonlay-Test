package model

import (
	"time"
)

type List struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:100"`
	Description string    `json:"description"`
	File        string    `json:"file" gorm:"size:300"`
	SubList     []SubList `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sub_list"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ListResponseRelational struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:100"`
	Description string    `json:"description"`
	File        string    `json:"file" gorm:"size:300"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ListRequest struct {
	Id          int    `json:"id" query:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
}

type ListCreateRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title" query:"title" validate:"required"`
	Description string `json:"description" query:"description" validate:"required"`
	FilePath    string `json:"file_path"`
}

type ListUpdateRequest struct {
	Id          int    `json:"id" query:"id" validate:"required"`
	Title       string `json:"title" query:"title" validate:"required"`
	Description string `json:"description" query:"description" validate:"required"`
	FilePath    string `json:"file" query:"file" gorm:"size:300"`
}

type ListDeleteRequest struct {
	Id int `json:"id" param:"id" query:"id" validate:"required"`
}
