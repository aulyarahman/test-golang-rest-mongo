package main

import (
	"github.com/aulyarahman/twitcat-service/app/model"
	"github.com/aulyarahman/twitcat-service/lib/config"
	"github.com/aulyarahman/twitcat-service/lib/db"
	"github.com/aulyarahman/twitcat-service/lib/logging"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

func init() {
	config.SetConfigFile("config", "lib/config", "json")
}

func main() {
	envConfig := getConfig()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{"*"},
		AllowCredentials: true,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(logging.MiddlewareLogging)


	//	Mongo
	mongo, err := db.Connect(envConfig.Mongo)
	if err != nil {
		log.Println(err)
		return
	}
}



func getConfig() model.EnvConfig {
	return model.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Mongo: db.MongoConfig{
			Timeout:  config.GetInt("database.mongodb.timeout"),
			DBname:   config.GetString("database.mongodb.dbname"),
			Username: config.GetString("database.mongodb.user"),
			Password: config.GetString("database.mongodb.password"),
			Host:     config.GetString("database.mongodb.host"),
			Port:     config.GetString("database.mongodb.port"),
		}
	}
}