package location

import (
	"context"
	"net/http"
	"reflect"

	"github.com/common-go/http"
	"github.com/common-go/model-builder"
	"github.com/common-go/search"
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
	idGenerator := builder.NewIdGenerator(generateId, false, false)
	modelBuilder := builder.NewModelBuilder(idGenerator.Generate, modelType, "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt", "userId")
	searchHandler := search.NewSearchHandler(locationService.Search, modelType, searchModelType, logError, nil)
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
