package app

import (
	"context"

	"github.com/common-go/mongo"
	. "github.com/common-go/service"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, mongoConfig mongo.MongoConfig) error {
	app, err := NewApp(ctx, mongoConfig)
	if err != nil {
		return err
	}

	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)

	userPath := "/users"
	r.HandleFunc(userPath, app.UserHandler.GetAll).Methods(GET)
	r.HandleFunc(userPath+"/search", app.UserHandler.Search).Methods(GET, POST)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Load).Methods(GET)
	r.HandleFunc(userPath, app.UserHandler.Insert).Methods(POST)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Update).Methods(PUT)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Patch).Methods(PATCH)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Delete).Methods(DELETE)

	locationPath := "/locations"
	r.HandleFunc(locationPath, app.LocationHandler.GetAll).Methods(GET)
	r.HandleFunc(locationPath+"/search", app.LocationHandler.Search).Methods(GET, POST)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Load).Methods(GET)
	r.HandleFunc(locationPath, app.LocationHandler.Create).Methods(POST)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Update).Methods(PUT)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Patch).Methods(PATCH)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Delete).Methods(DELETE)
	return nil
}
