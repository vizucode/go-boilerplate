package main

import (
	"os"

	flightproviders "gitlab.com/tixia-backend/go-boilerplate/apps/repositories/flight_providers"
	"gitlab.com/tixia-backend/go-boilerplate/apps/repositories/psql"
	"gitlab.com/tixia-backend/go-boilerplate/apps/router/cmd"
	"gitlab.com/tixia-backend/go-boilerplate/apps/service/transaction"
	errorhandler "gitlab.com/tixia-backend/go-boilerplate/helpers/error_handler"
	"k8s.io/utils/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vizucode/gokit/adapter/dbc"
	"github.com/vizucode/gokit/utils/constant"
)

func main() {

	var serviceName = "CronJob"

	/*
		Library
	*/
	gormUserConn := dbc.NewGormConnection(
		dbc.SetGormURIConnection(env.GetString("DB_TRANSACTION_DSN", "")),
		dbc.SetGormDriver(constant.Postgres),
		dbc.SetGormMaxIdleConnection(2),
		dbc.SetGormMaxPoolConnection(50),
		dbc.SetGormMinPoolConnection(10),
		dbc.SetGormSkipTransaction(true),
		dbc.SetGormServiceName(serviceName),
	)

	/*
		Repositories
	*/
	postgreDB := psql.NewPsql(gormUserConn.DB)
	flightProvider := flightproviders.NewFlightProviders(
		env.GetString("RPC_CORE_SERVICE_HOST", ""),
		env.GetString("RPC_CORE_SERVICE_PORT", ""),
	)

	/*
		Service
	*/
	transactions := transaction.NewTransaction(
		postgreDB,
		flightProvider,
	)

	/*
		REST server
	*/

	cmdHandler := cmd.NewCmd(transactions)

	cmdHandler.DoExpired()

	app := fiber.New(fiber.Config{
		ErrorHandler: errorhandler.FiberErrHandler,
	})

	app.Use(recover.New())

	app.Listen(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
