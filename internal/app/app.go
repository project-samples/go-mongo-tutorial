package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/common-go/health"
	"github.com/common-go/log"
	"github.com/common-go/mongo"
	"github.com/common-go/search"
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
type TestSE struct {
	PageIndex     int64 `mapstructure:"page_index" json:"pageIndex,omitempty" gorm:"column:pageindex" bson:"pageIndex,omitempty" dynamodbav:"pageIndex,omitempty" firestore:"pageIndex,omitempty"`
	Fields        []string                 `mapstructure:"fields" json:"fields,omitempty" gorm:"column:fields" bson:"fields,omitempty" dynamodbav:"fields,omitempty" firestore:"fields,omitempty"`
	Sort          string                   `mapstructure:"sort" json:"sort,omitempty" gorm:"column:sortfield" bson:"sort,omitempty" dynamodbav:"sort,omitempty" firestore:"sort,omitempty"`
	RefId         string                   `mapstructure:"refid" json:"refId,omitempty" gorm:"column:refid" bson:"refId,omitempty" dynamodbav:"refId,omitempty" firestore:"refId,omitempty"`

	PageSize      int64 `mapstructure:"page_size" json:"pageSize,omitempty" gorm:"column:pagesize" bson:"pageSize,omitempty" dynamodbav:"pageSize,omitempty" firestore:"pageSize,omitempty"`
	FirstPageSize int64 `mapstructure:"first_page_size" json:"firstPageSize,omitempty" gorm:"column:firstpagesize" bson:"firstPageSize,omitempty" dynamodbav:"firstPageSize,omitempty" firestore:"firstPageSize,omitempty"`

	Page          int64                    `mapstructure:"page" json:"page,omitempty" gorm:"column:pageindex" bson:"page,omitempty" dynamodbav:"page,omitempty" firestore:"page,omitempty"`
	Limit         int64                    `mapstructure:"limit" json:"limit,omitempty" gorm:"column:limit" bson:"limit,omitempty" dynamodbav:"limit,omitempty" firestore:"limit,omitempty"`
	FirstLimit    int64                    `mapstructure:"first_limit" json:"firstLimit,omitempty" gorm:"column:firstlimit" bson:"firstLimit,omitempty" dynamodbav:"firstLimit,omitempty" firestore:"firstLimit,omitempty"`

	CurrentUserId string                   `mapstructure:"current_user_id" json:"currentUserId,omitempty" gorm:"column:currentuserid" bson:"currentUserId,omitempty" dynamodbav:"currentUserId,omitempty" firestore:"currentUserId,omitempty"`
	Keyword       string                   `mapstructure:"keyword" json:"keyword,omitempty" gorm:"column:keyword" bson:"keyword,omitempty" dynamodbav:"keyword,omitempty" firestore:"keyword,omitempty"`
	Excluding     map[string][]interface{} `mapstructure:"excluding" json:"excluding,omitempty" gorm:"column:excluding" bson:"excluding,omitempty" dynamodbav:"excluding,omitempty" firestore:"excluding,omitempty"`
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
	s := TestSE{
		Sort: "TestSort",
		RefId: "Test2",
		Fields: []string{"field1", "field2"},
	}
	s1b, s2b, s3b := search.GetFieldsAndSortAndRefId(&s)
	s1, s2, s3 := search.GetFieldsAndSortAndRefId(s)

	x := fmt.Sprintf("%v %s %s %v %s %s", s1, s2, s3, s1b, s2b, s3b)
	fmt.Println(x)
	logError := log.ErrorMsg

	validator := validator.NewDefaultValidator()
	userService := services.NewUserService(db, search.GetSort)
	userHandler := handlers.NewUserHandler(userService, validator, logError)

	locationService := location.NewLocationService(db, search.GetSort)
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
