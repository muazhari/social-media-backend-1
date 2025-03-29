package configs

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
	"social-media-backend-1/internal/inners/models/entities"
)

type TwoDatabaseConfig struct {
	Connection  *sql.DB
	AccountData []*entities.Account
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
		AccountData: []*entities.Account{
			{
				ID:               uuid.New(),
				Email:            "zero@mail.com",
				Password:         "zero",
				Name:             "zero",
				TotalPostLike:    2,
				TotalChatMessage: 2,
			},
			{
				ID:               uuid.New(),
				Email:            "one@mail.com",
				Password:         "one",
				Name:             "one",
				TotalPostLike:    1,
				TotalChatMessage: 1,
			},
		},
	}

	return twoDatabaseConfig
}
