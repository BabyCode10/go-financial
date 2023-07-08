package services

import (
	"database/sql"
	"errors"
	"go-financial/app/models"
	"go-financial/app/requests"
	"time"

	"github.com/jmoiron/sqlx"
)

type SDatabaseTransaction struct {
	conn *sqlx.DB
}

func Transaction(conn *sqlx.DB) *SDatabaseTransaction {
	return &SDatabaseTransaction{
		conn: conn,
	}
}

func (databaseTransaction *SDatabaseTransaction) GetByUser(userId string) (*[]models.STransaction, error) {
	query := `
		SELECT
			*
		FROM
			transactions
		WHERE
			user_id = $1
	`

	var transactions []models.STransaction

	err := databaseTransaction.conn.Select(
		&transactions,
		query,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (databaseTransaction *SDatabaseTransaction) Store(userId string, request *requests.STransactionStoreRequest) (*models.STransaction, error) {
	query := `
		INSERT INTO
			transactions(user_id, category_id, type, currency, note, amount)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	var transaction models.STransaction

	err := databaseTransaction.conn.Get(
		&transaction,
		query,
		userId,
		request.CategoryId,
		request.Type,
		request.Currency,
		request.Note,
		request.Amount,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (databaseTransaction *SDatabaseTransaction) FindByUser(userId string, transactionId string) (*models.STransaction, error) {
	query := `
		SELECT
			*
		FROM
			transactions
		WHERE
			id = $1
		AND
			user_id = $2
	`

	var transaction models.STransaction

	err := databaseTransaction.conn.Get(
		&transaction,
		query,
		transactionId,
		userId,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("Transaction not found.")
	}

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (databaseTransaction *SDatabaseTransaction) Update(userId string, transactionId string, request *requests.STransactionUpdateRequest) (*models.STransaction, error) {
	query := `
		UPDATE
			transactions
		SET
			category_id = $1,
			type = $2,
			currency = $3,
			note = $4,
			amount = $5,
			updated_at = $6
		WHERE
			id = $7
		AND
			user_id = $8
		RETURNING *
	`

	var transaction models.STransaction

	err := databaseTransaction.conn.Get(
		&transaction,
		query,
		request.CategoryId,
		request.Type,
		request.Currency,
		request.Note,
		request.Amount,
		time.Now(),
		transactionId,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (databaseTransaction *SDatabaseTransaction) Delete(userId string, transactionId string) (*models.STransaction, error) {
	query := `
		DELETE FROM
			transactions
		WHERE
			id = $1
		AND
			user_id = $2
		RETURNING *
	`

	var transaction models.STransaction

	err := databaseTransaction.conn.Get(
		&transaction,
		query,
		transactionId,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
