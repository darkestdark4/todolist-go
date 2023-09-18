package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
)

func ToTodoResponse(todo domain.Todo) web.TodoResponse {
	return web.TodoResponse{
		Id:          todo.Id,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
	}
}

func ToTodoResponses(todos []domain.Todo) []web.TodoResponse {
	var todoResponses []web.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, ToTodoResponse(todo))
	}

	return todoResponses
}
