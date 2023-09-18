package service

import (
	"context"
	"golang-restful-api/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.TodoResponse
	FindAll(ctx context.Context) []web.TodoResponse
}
