package app

import (
	"github.com/common-go/http"
	"github.com/common-go/log"
	m "github.com/common-go/middleware"
	"github.com/common-go/mongo"
)

type Root struct {
	Server     server.ServerConfig `mapstructure:"server"`
	Mongo      mongo.MongoConfig   `mapstructure:"mongo"`
	Log        log.Config          `mapstructure:"log"`
	MiddleWare m.LogConfig         `mapstructure:"middleware"`
}
