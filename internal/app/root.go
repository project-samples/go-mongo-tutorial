package app

type Root struct {
	Server   ServerConfig `mapstructure:"server"`
	Mongo    MongoConfig  `mapstructure:"mongo"`
	Response string       `mapstructure:"response"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type MongoConfig struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}
