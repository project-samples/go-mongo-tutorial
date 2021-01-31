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

	return nil
}
