package models

import (
	"go-financial/app/requests"
	"time"
)

type SCategory struct {
	Id        int64
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type ICategory interface {
	Get() (*[]SCategory, error)
	Store(request *requests.SCategoryStoreRequest) (*SCategory, error)
	Find(categoryId string) (*SCategory, error)
	Update(categoryId string, request *requests.SCategoryUpdateRequest) (*SCategory, error)
	Delete(categoryId string) (*SCategory, error)
}
