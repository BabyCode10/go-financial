package models

import (
	"go-financial/app/requests"
	"time"
)

type SAuth struct {
	Id        int64
	Username  string     `db:"username"`
	Name      string     `db:"name"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type IAuth interface {
	Store(request *requests.SAuthRegisterRequest) (*SAuth, error)
	FindByUsername(username string) (*SAuth, error)
	FindById(userId string) (*SAuth, error)
	Update(userId string, request *requests.SAuthUpdateRequest) (*SAuth, error)
}
