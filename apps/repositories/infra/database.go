package infra

import (
	"net/http"

	"gitlab.com/tixia-backend/go-boilerplate/apps/models"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/logger"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/rpcstd"
	errorhandler "gitlab.com/tixia-backend/go-boilerplate/helpers/error_handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogError(err)
		return db, errorhandler.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, "database can't connected, please check")
	}

	// migration
	err = db.AutoMigrate(
		&models.TodoList{},
	)
	if err != nil {
		logger.LogError(err)
		return db, errorhandler.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, "migration error")
	}

	return db, err
}
