package exception

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err interface{}) {
	if notFoundError(w, req, err) {
		return
	}

	if validationError(w, req, err) {
		return
	}

	StatusInternalServerError(w, req, err)
}

func notFoundError(w http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Resource Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func StatusInternalServerError(w http.ResponseWriter, req *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
