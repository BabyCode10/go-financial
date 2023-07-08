package models

import (
	"go-financial/app/requests"
	"time"
)

type STransaction struct {
	Id         int64
	UserId     string     `db:"user_id"`
	CategoryId string     `db:"category_id"`
	Type       string     `db:"type"`
	Currency   string     `db:"currency"`
	Note       string     `db:"note"`
	Amount     int32      `db:"amount"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}

type ITransaction interface {
	GetByUser(userId string) (*[]STransaction, error)
	Store(userId string, request *requests.STransactionStoreRequest) (*STransaction, error)
	FindByUser(userId string, transactionId string) (*STransaction, error)
	Update(userId string, transactionId string, request *requests.STransactionUpdateRequest) (*STransaction, error)
	Delete(userId string, transactionId string) (*STransaction, error)
}
