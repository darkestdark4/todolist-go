package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repository *TodoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "INSERT INTO todolist(name, description, status) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, todo.Name, todo.Description, todo.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todo.Id = int(id)
	return todo
}

func (repository *TodoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "UPDATE todolist SET name = ?, description = ?, status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, todo.Name, todo.Description, todo.Status, todo.Id)
	helper.PanicIfError(err)

	return todo
}

func (repository *TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) {
	SQL := "DELETE FROM todolist WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, todo.Id)
	helper.PanicIfError(err)
}

func (repository *TodoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error) {
	SQL := "SELECT * FROM todolist WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, todoId)
	helper.PanicIfError(err)
	defer rows.Close()

	todo := domain.Todo{}
	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.Status)
		helper.PanicIfError(err)

		return todo, nil
	} else {
		return todo, errors.New("todo list is not found")
	}
}

func (repository *TodoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo {
	SQL := "SELECT * from todolist"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []domain.Todo
	for rows.Next() {
		todo := domain.Todo{}
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.Status)
		helper.PanicIfError(err)

		todos = append(todos, todo)
	}

	return todos
}
