package services

import (
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"

	m "github.com/common-go/mongo"
	"github.com/common-go/search"
	"github.com/common-go/service"

	. "go-service/internal/models"
)

type MongoUserService struct {
	search.SearchService
	service.GenericService
}

func NewUserService(db *mongo.Database) *MongoUserService {
	var model User
	collectionName := "users"
	modelType := reflect.TypeOf(model)
	queryBuilder := m.NewQueryBuilder(modelType)
	searcher, writer := m.NewSearchWriterWithQuery(db, collectionName, modelType, queryBuilder.BuildQuery)
	return &MongoUserService{SearchService: searcher, GenericService: writer}
}
