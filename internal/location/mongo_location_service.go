package location

import (
	m "github.com/common-go/mongo"
	"github.com/common-go/search"
	"github.com/common-go/service"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type MongoLocationService struct {
	service.GenericService
	search.SearchService
	LocationMapper m.Mapper
}

func NewMongoLocationService(db *mongo.Database) *MongoLocationService {
	var model Location
	modelType := reflect.TypeOf(model)
	mapper := m.NewMapper(modelType)
	queryBuilder := m.NewQueryBuilder(modelType)
	genericService, searchService := m.NewSearchWriterWithQuery(db, "location", modelType, queryBuilder.BuildQuery, mapper)
	return &MongoLocationService{genericService, searchService, mapper}
}
