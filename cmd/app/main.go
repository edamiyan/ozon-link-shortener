package main

import (
	"flag"
	handler "github.com/edamiyan/ozon-link-shortener/internal/handler/v1"
	inMemory "github.com/edamiyan/ozon-link-shortener/internal/repository/in_memory"
	"github.com/edamiyan/ozon-link-shortener/internal/repository/postgres"
	"github.com/edamiyan/ozon-link-shortener/internal/service/link"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	databaseURLKey = "DATABASE_URL"
	portKey        = "PORT"
)

func main() {
	var dbFlag bool
	flag.BoolVar(&dbFlag, "db", false, "Run with DB postgres:")
	flag.Parse()
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %v", err)
	}

	var repos link.Repository
	if dbFlag {
		dbConfig := os.Getenv(databaseURLKey)
		if dbConfig == "" {
			logrus.Info("empty env config")
			dbConfig = viper.GetString(databaseURLKey)
		}

		db, err := postgres.NewPostgresDB(dbConfig)
		if err != nil {
			logrus.Fatalf("failed to initialize db: %v", err)
		}
		repos = postgres.NewRepository(db)
	} else {
		const mapLen = 64
		repos = inMemory.NewRepository(mapLen)
	}

	linkService := link.NewService(repos)
	handlers := handler.NewHandler(linkService)

	app := echo.New()
	handlers.Init(app)
	port := viper.GetString(portKey)

	if err := app.Start(":" + port); err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
