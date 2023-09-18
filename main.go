package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	repository := repository.NewTodoRepository()
	service := service.NewTodoService(repository, db, validate)
	todoController := controller.NewTodoController(service)

	router := httprouter.New()

	router.GET("/api/todo", todoController.FindAll)
	router.GET("/api/todo/:id", todoController.FindById)
	router.POST("/api/todo", todoController.Create)
	router.PUT("/api/todo/:id", todoController.Update)
	router.DELETE("/api/todo/:id", todoController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
