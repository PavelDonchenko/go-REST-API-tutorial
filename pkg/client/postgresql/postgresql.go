package postgresql

import (
	"context"
	"fmt"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/config"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttemps int, sc config.StorageConfig) (con *pgx.Conn, err error) {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		con, err = pgx.Connect(ctx, dsn)
		if err != nil {
			fmt.Println("failed to connect to progresql")
			return err
		}
		return nil
	}, maxAttemps, 5*time.Second)
	if err != nil {
		log.Fatalln("error do with tries postgresql")
	}
	return con, nil
}
