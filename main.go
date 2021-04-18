package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/common-go/config"
	"github.com/common-go/log"
	m "github.com/common-go/middleware"
	sv "github.com/common-go/service"
	"github.com/gorilla/mux"

	"go-service/internal/app"
)

func main() {
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}

	r := mux.NewRouter()
	log.Initialize(conf.Log)
	r.Use(m.BuildContext)
	logger := m.NewStructuredLogger()
	r.Use(m.Logger(conf.MiddleWare, log.InfoFields, logger))
	r.Use(m.Recover(log.ErrorMsg))

	er2 := app.Route(r, context.Background(), conf.Mongo)
	if er2 != nil {
		panic(er2)
	}
	fmt.Println(sv.GetStartMessage(conf.Server))
	http.ListenAndServe(sv.GetServerAddress(conf.Server.Port), r)
}
