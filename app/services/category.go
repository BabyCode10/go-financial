package services

import (
	"database/sql"
	"errors"
	"go-financial/app/models"
	"go-financial/app/requests"
	"time"

	"github.com/jmoiron/sqlx"
)

type SDatabaseCategory struct {
	conn *sqlx.DB
}

func Category(conn *sqlx.DB) *SDatabaseCategory {
	return &SDatabaseCategory{
		conn: conn,
	}
}

func (databaseCategory *SDatabaseCategory) Get() (*[]models.SCategory, error) {
	query := `
		SELECT
			*
		FROM
			categories
	`

	var categories []models.SCategory

	err := databaseCategory.conn.Select(
		&categories,
		query,
	)

	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (databaseCategory *SDatabaseCategory) Store(request *requests.SCategoryStoreRequest) (*models.SCategory, error) {
	query := `
		INSERT INTO 
			categories(name)
		VALUES 
			($1)
		RETURNING *
	`

	var category models.SCategory

	err := databaseCategory.conn.Get(
		&category,
		query,
		request.Name,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (databaseCategory *SDatabaseCategory) Find(categoryId string) (*models.SCategory, error) {
	query := `
		SELECT 
			*
		FROM
			categories
		WHERE
			id = $1
		LIMIT 1
	`

	var category models.SCategory

	err := databaseCategory.conn.Get(
		&category,
		query,
		categoryId,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("User not found.")
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (databaseCategory *SDatabaseCategory) Update(categoryId string, request *requests.SCategoryUpdateRequest) (*models.SCategory, error) {
	query := `
		UPDATE
			categories
		SET
			name = $1,
			updated_at = $2
		WHERE
			id = $3
		RETURNING *
	`

	var category models.SCategory

	err := databaseCategory.conn.Get(
		&category,
		query,
		request.Name,
		time.Now(),
		categoryId,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (databaseCategory *SDatabaseCategory) Delete(categoryId string) (*models.SCategory, error) {
	query := `
		DELETE FROM
			categories
		WHERE
			id = $1
		RETURNING *
	`

	var category models.SCategory

	err := databaseCategory.conn.Get(
		&category,
		query,
		categoryId,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
