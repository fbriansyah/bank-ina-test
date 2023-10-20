package main

import (
	db "github.com/fbriansyah/bank-ina-test/internal/adapter/database"
	"github.com/fbriansyah/bank-ina-test/internal/adapter/gin"
	"github.com/fbriansyah/bank-ina-test/internal/adapter/token"
	"github.com/fbriansyah/bank-ina-test/internal/application"
	"github.com/fbriansyah/bank-ina-test/util"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %v\n", err)
	}

	sqlDB := connectToDB(config.DBDriver, config.DBSource)
	if sqlDB == nil {
		log.Fatal().Msgf("cannot connect to db:%v\n", err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	databaseAdapter := db.NewDatabaseAdapter(sqlDB)
	tokenAdapter, err := token.NewPasetoMaker(config.TokenSymmetric)
	if err != nil {
		log.Fatal().Msgf("cannot connect to db:%v\n", err)
	}

	service := application.NewService(databaseAdapter, tokenAdapter)

	serverAdapter := gin.NewAdapter(service)

	serverAdapter.Start(config.HTTPServerAddress)
}
