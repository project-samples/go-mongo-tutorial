package app

import (
	"context"
	"strings"

	"github.com/common-go/health"
	"github.com/common-go/log"
	"github.com/common-go/mongo"
	"github.com/common-go/validator"
	"github.com/google/uuid"

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
	HealthHandler   *health.HealthHandler
	UserHandler     *handlers.UserHandler
	LocationHandler *location.LocationHandler
}

func NewApp(ctx context.Context, mongoConfig mongo.MongoConfig) (*ApplicationContext, error) {
	db, err := mongo.SetupMongo(ctx, mongoConfig)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg

	validator := validator.NewDefaultValidator()
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService, generateId, validator, logError)


	locationService := location.NewLocationService(db)
	locationHandler := location.NewLocationHandler(locationService, generateId, validator, logError)

	mongoChecker := mongo.NewHealthChecker(db)
	checkers := []health.HealthChecker{mongoChecker}
	healthHandler := health.NewHealthHandler(checkers)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:     userHandler,
		LocationHandler: locationHandler,
	}, nil
}
