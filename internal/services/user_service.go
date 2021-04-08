package services

import (
	"github.com/common-go/search"
	"github.com/common-go/service"
)

type UserService interface {
	search.SearchService
	service.GenericService
}
