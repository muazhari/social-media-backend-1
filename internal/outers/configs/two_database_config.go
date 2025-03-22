package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

type TwoDatabaseConfig struct {
	Connection *sql.DB
}

func NewTwoDatabaseConfig() *TwoDatabaseConfig {
	connection, openErr := sql.Open(
		"pgx",
		fmt.Sprintf(
			"postgresql://%s@%s:%s/%s",
			os.Getenv("DATASTORE_2_HOST"),
			os.Getenv("DATASTORE_2_PORT"),
			os.Getenv("DATASTORE_2_USER"),
			os.Getenv("DATASTORE_2_PASSWORD"),
			os.Getenv("DATASTORE_2_DBNAME"),
		),
	)
	if openErr != nil {
		panic(openErr)
	}

	twoDatabaseConfig := &TwoDatabaseConfig{
		Connection: connection,
	}

	return twoDatabaseConfig
}
