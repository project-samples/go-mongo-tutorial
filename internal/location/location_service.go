package location

import (
	"github.com/common-go/search"
	"github.com/common-go/service"
)

type LocationService interface {
	search.SearchService
	service.GenericService
}
