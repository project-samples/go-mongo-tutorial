package app

import (
	"context"
	"strings"

	"github.com/common-go/log"
	"github.com/common-go/validator"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/internal/handlers"
	"go-service/internal/location"
	"go-service/internal/services"
)

func randomId() string {
	id := uuid.New()
	return strings.Replace(id.String(), "-", "", -1)
}
func generateId(ctx context.Context) (string, error) {
	id := randomId()
	return id, nil
}
type ApplicationContext struct {
	UserHandler     *handlers.UserHandler
	LocationHandler *location.LocationHandler
}

func NewApp(context context.Context, mongoConfig MongoConfig) (*ApplicationContext, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(mongoConfig.Uri))
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg
	db := client.Database(mongoConfig.Database)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	validator := validator.NewDefaultValidator()
	locationService := location.NewMongoLocationService(db)
	locationHandler := location.NewLocationHandler(locationService, generateId, validator, logError)

	return &ApplicationContext{
		UserHandler: userHandler,
		LocationHandler: locationHandler,
	}, nil
}
