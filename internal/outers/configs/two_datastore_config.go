package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type TwoDatastoreConfig struct {
	Connection *sql.DB
}

func NewTwoDatastoreConfig() *TwoDatastoreConfig {
	connection, openErr := sql.Open(
		"pgx",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s",
			os.Getenv("DATASTORE_2_USER"),
			os.Getenv("DATASTORE_2_PASSWORD"),
			os.Getenv("DATASTORE_2_HOST"),
			os.Getenv("DATASTORE_2_PORT"),
			os.Getenv("DATASTORE_2_DATABASE"),
		),
	)
	if openErr != nil {
		panic(openErr)
	}

	twoDatabaseConfig := &TwoDatastoreConfig{
		Connection: connection,
	}

	return twoDatabaseConfig
}
