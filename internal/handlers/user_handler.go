package handlers

import (
	"context"
	"net/http"
	"reflect"

	"github.com/common-go/http"
	"github.com/common-go/search"
	"github.com/common-go/service"
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

func NewUserHandler(userService UserService, generateId func(context.Context) (string, error), validator validator.Validator, logError func(context.Context, string)) *UserHandler {
	modelType := reflect.TypeOf(User{})
	searchModelType := reflect.TypeOf(UserSM{})
	idGenerator := service.NewIdGenerator(generateId, false, false)
	modelBuilder := service.NewModelBuilder(idGenerator.Generate, modelType, "", "userId", "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt")
	searchHandler := search.NewSearchHandler(userService.Search, searchModelType, logError, nil)
	genericHandler := server.NewGenericHandler(userService, modelType, modelBuilder, logError, validator.Validate)
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
