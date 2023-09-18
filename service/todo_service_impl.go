package service

import (
	"context"
	"database/sql"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, DB *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{
		Name:        request.Name,
		Description: request.Description,
		Status:      request.Status,
	}

	todo = service.TodoRepository.Save(ctx, tx, todo)

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todo.Name = request.Name
	todo.Description = request.Description
	todo.Status = request.Status

	todo = service.TodoRepository.Update(ctx, tx, todo)

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoRepository.Delete(ctx, tx, todo)
}

func (service *TodoServiceImpl) FindById(ctx context.Context, todoId int) web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) FindAll(ctx context.Context) []web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.TodoRepository.FindAll(ctx, tx)

	return helper.ToTodoResponses(categories)
}
