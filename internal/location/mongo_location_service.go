package location

import (
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"

	m "github.com/common-go/mongo"
	"github.com/common-go/search"
	"github.com/common-go/service"
)

type MongoLocationService struct {
	search.SearchService
	service.GenericService
	Mapper m.Mapper
}

func NewLocationService(db *mongo.Database) *MongoLocationService {
	var model Location
	modelType := reflect.TypeOf(model)
	mapper := m.NewMapper(modelType)
	queryBuilder := m.NewQueryBuilder(modelType)
	searchService, genericService := m.NewSearchWriterWithQuery(db, "location", modelType, queryBuilder.BuildQuery, mapper)
	return &MongoLocationService{SearchService: searchService, GenericService: genericService, Mapper: mapper}
}
