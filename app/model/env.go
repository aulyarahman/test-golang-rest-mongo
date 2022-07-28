package model

import "github.com/aulyarahman/twitcat-service/lib/db"

type (
	EnvConfig struct {
		Host  string
		Port  int
		Mongo db.MongoConfig
	}
)
