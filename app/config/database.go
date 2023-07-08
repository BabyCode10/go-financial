package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Database() (*sqlx.DB, error) {
	var config Config

	config, err := Load()

	if err != nil {
		panic("Failed to load config.")
	}

	info := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBDriver,
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase,
	)

	conn, err := sqlx.Open(config.DBDriver, info)

	if err != nil {
		message := fmt.Errorf("Database failed to connect: %w", err)

		panic(message)
	}

	err = conn.Ping()

	if err != nil {
		message := fmt.Errorf("Database failed to ping: %w", err)

		panic(message)
	}

	return conn, err
}
