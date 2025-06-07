package psql

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gitlab.com/tixia-backend/go-boilerplate/apps/models"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/httpstd"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/rpcstd"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/statuses"
	"gorm.io/gorm"
)

func (psql *psql) GetTransactionByCode(ctx context.Context) (resp []models.TbTransaction, err error) {
	dx := psql.transaction.WithContext(ctx).
		Where("expired_at >= ?", time.Now().Add(7*time.Hour+10*time.Minute)).
		Where("status", statuses.BOOKED).
		Or("status", statuses.INITATED).
		Preload("FlightTransaction.FlightJourney").
		Preload("Passenger").
		Preload("Insurance")

	err = dx.Find(&resp).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, httpstd.InternalServerError)
	}

	return resp, nil
}
