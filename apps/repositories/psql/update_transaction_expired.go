package psql

import (
	"context"
	"net/http"
	"time"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gitlab.com/tixia-backend/go-boilerplate/apps/models"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/statuses"
)

func (psql *psql) UpdateTransactionToExpired(ctx context.Context) (err error) {

	var resultTransactions []models.TbTransaction
	var expiredAt time.Time = time.Now().Add(time.Hour * 7).Add(time.Minute * 10)

	psql.transaction.Model(&models.TbTransaction{}).Debug().
		Where("expired_at < ?", expiredAt).
		Where("status IN ?", []string{statuses.INITATED, statuses.BOOKED}).
		Find(&resultTransactions)

	err = psql.transaction.Model(&models.TbTransaction{}).
		Select("status").
		Where("expired_at < ?", expiredAt).
		Where("status IN ?", []string{statuses.INITATED, statuses.BOOKED}).
		Update("status", statuses.EXPIRED).Error

	if err != nil {
		logger.Log.Debug(ctx, err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, "00", err.Error())
	}

	for _, transaction := range resultTransactions {
		err = psql.transaction.Model(&models.TbFlightTransaction{}).
			Where("transaction_id = ?", transaction.Id).
			Update("status", statuses.EXPIRED).Error
		if err != nil {
			logger.Log.Debug(ctx, err)
		}

		err = psql.transaction.Model(&models.TbHotelTransaction{}).
			Where("transaction_id = ?", transaction.Id).
			Update("status", statuses.EXPIRED).Error
		if err != nil {
			logger.Log.Debug(ctx, err)
		}
	}

	return nil
}
