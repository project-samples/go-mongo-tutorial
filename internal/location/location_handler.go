package location

import (
	"context"
	"net/http"
	"reflect"

	"github.com/common-go/http"
	"github.com/common-go/search"
	"github.com/common-go/service"
	"github.com/common-go/validator"
)

type LocationHandler struct {
	*server.GenericHandler
	*search.SearchHandler
	Service LocationService
}

func NewLocationHandler(locationService LocationService, generateId func(context.Context) (string, error), validator validator.Validator, logError func(context.Context, string)) *LocationHandler {
	modelType := reflect.TypeOf(Location{})
	searchModelType := reflect.TypeOf(LocationSM{})
	idGenerator := service.NewIdGenerator(generateId, false, false)
	modelBuilder := service.NewModelBuilder(idGenerator.Generate, modelType, "", "userId", "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt")
	searchHandler := search.NewSearchHandler(locationService.Search, searchModelType, logError, nil)
	genericHandler := server.NewGenericHandler(locationService, modelType, modelBuilder, logError, validator.Validate)
	return &LocationHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: locationService}
}

func (h *LocationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	server.Respond(w, r, http.StatusOK, result)
}
