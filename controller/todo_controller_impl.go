package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (controller *TodoControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todoCreateRequest := web.TodoCreateRequest{}
	helper.ReadFromRequestBody(r, &todoCreateRequest)

	todoResponse := controller.TodoService.Create(r.Context(), todoCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todoUpdateRequest := web.TodoUpdateRequest{}
	helper.ReadFromRequestBody(r, &todoUpdateRequest)

	todoId := p.ByName("id")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoUpdateRequest.Id = id

	todoResponse := controller.TodoService.Update(r.Context(), todoUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todoId := p.ByName("id")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	controller.TodoService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todoId := p.ByName("id")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoResponse := controller.TodoService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todoResponses := controller.TodoService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
