package handlers

import (
	"context"
	"net/http"
	"reflect"

	"github.com/common-go/http"
	"github.com/common-go/search"
	"github.com/common-go/validator"

	. "go-service/internal/models"
	. "go-service/internal/search-models"
	. "go-service/internal/services"
)

type UserHandler struct {
	*server.GenericHandler
	*search.SearchHandler
	Service UserService
}

func NewUserHandler(userService UserService, validator validator.Validator, logError func(context.Context, string)) *UserHandler {
	modelType := reflect.TypeOf(User{})
	searchModelType := reflect.TypeOf(UserSM{})
	searchHandler := search.NewSearchHandler(userService.Search, modelType, searchModelType, logError, nil)
	genericHandler := server.NewGenericHandler(userService, modelType, nil, logError, validator.Validate)
	return &UserHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: userService}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	server.Respond(w, r, http.StatusOK, result)
}
