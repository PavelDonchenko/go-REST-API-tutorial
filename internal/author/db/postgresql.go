package db

import (
	"context"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/author"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/pkg/client/postgresql"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger logging.Logger
}

func (r *repository) Create(ctx context.Context, author author.Author) (string, error) {
	//q := `INSERT INTO author (name) VALUES ($1)`
	panic("implement me")
}

func (r *repository) FindAll(ctx context.Context) ([]author.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, author author.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
