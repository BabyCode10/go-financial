package services

import (
	"go-financial/app/models"
	"go-financial/app/requests"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type SDatabaseAuth struct {
	conn *sqlx.DB
}

func User(conn *sqlx.DB) *SDatabaseAuth {
	return &SDatabaseAuth{
		conn: conn,
	}
}

func (databaseAuth *SDatabaseAuth) Store(request *requests.SAuthRegisterRequest) (*models.SAuth, error) {
	query := `
		INSERT INTO
			users(username, name, password)
		VALUES
			($1, $2, $3)
		RETURNING *
	`

	var user models.SAuth

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	err = databaseAuth.conn.Get(
		&user,
		query,
		request.Username,
		request.Name,
		hash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (databaseAuth *SDatabaseAuth) FindByUsername(username string) (*models.SAuth, error) {
	query := `
		SELECT
			*
		FROM
			users
		WHERE
			username = $1
		LIMIT 1
	`

	var user models.SAuth

	err := databaseAuth.conn.Get(
		&user,
		query,
		username,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (databaseAuth *SDatabaseAuth) FindById(userId string) (*models.SAuth, error) {
	query := `
		SELECT
			*
		FROM
			users
		WHERE
			id = $1
		LIMIT 1
	`

	var user models.SAuth

	err := databaseAuth.conn.Get(
		&user,
		query,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (databaseAuth *SDatabaseAuth) Update(username string, request *requests.SAuthUpdateRequest) (*models.SAuth, error) {
	query := `
		UPDATE
			users
		SET
			name = $1,
			password = $2,
			updated_at = $3
		WHERE
			username = $4
		RETURNING *
	`

	var user models.SAuth

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	err = databaseAuth.conn.Get(
		&user,
		query,
		request.Name,
		hash,
		time.Now(),
		username,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
