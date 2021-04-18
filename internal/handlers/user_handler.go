package handlers

import (
	"context"
	"net/http"
	"reflect"

	"github.com/common-go/search"
	sv "github.com/common-go/service"

	. "go-service/internal/models"
	. "go-service/internal/search-models"
	. "go-service/internal/services"
)

type UserHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	Service UserService
}

func NewUserHandler(userService UserService, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string)) *UserHandler {
	modelType := reflect.TypeOf(User{})
	searchModelType := reflect.TypeOf(UserSM{})
	searchHandler := search.NewSearchHandler(userService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewGenericHandler(userService, modelType, nil, logError, validate)
	return &UserHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: userService}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
