package app

import (
	"context"

	"github.com/go-chi/chi"
)

func Route(r *chi.Mux, context context.Context, mongoConfig MongoConfig) error {
	app, err := NewApp(context, mongoConfig)
	if err != nil {
		return err
	}

	userPath := "/users"
	r.Get(userPath, app.UserHandler.GetAll)
	r.Get(userPath+ "/{id}", app.UserHandler.Load)
	r.Post(userPath, app.UserHandler.Insert)
	r.Put(userPath+ "/{id}", app.UserHandler.Update)
	r.Delete(userPath+ "/{id}", app.UserHandler.Delete)

	locationPath := "/locations"
	r.Get(locationPath, app.LocationHandler.GetAll)
	// r.Get(locationPath, app.LocationHandler.Search)
	r.Post(locationPath + "/search", app.LocationHandler.Search)
	r.Get(locationPath+ "/{id}", app.LocationHandler.Load)
	r.Post(locationPath, app.LocationHandler.Insert)
	r.Put(locationPath+ "/{id}", app.LocationHandler.Update)
	r.Delete(locationPath+ "/{id}", app.LocationHandler.Delete)
	return nil
}
