package location

import (
	"context"
	"net/http"
	"reflect"

	"github.com/common-go/model-builder"
	"github.com/common-go/search"
	sv "github.com/common-go/service"
)

type LocationHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	Service LocationService
}

func NewLocationHandler(locationService LocationService, generateId func(context.Context) (string, error), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string)) *LocationHandler {
	modelType := reflect.TypeOf(Location{})
	searchModelType := reflect.TypeOf(LocationSM{})
	idGenerator := builder.NewIdGenerator(generateId, false, false)
	modelBuilder := builder.NewModelBuilder(idGenerator.Generate, modelType, "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt", "userId")
	searchHandler := search.NewSearchHandler(locationService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewGenericHandler(locationService, modelType, modelBuilder, logError, validate)
	return &LocationHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: locationService}
}

func (h *LocationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
