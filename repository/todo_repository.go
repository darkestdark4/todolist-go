package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/domain"
)

type TodoRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Todo, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo
}
