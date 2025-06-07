package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/vizucode/gokit/utils/env"
	"gitlab.com/tixia-backend/go-boilerplate/apps/models"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/logger"
)

func (uc *transaction) DoExpired(ctx context.Context) (err error) {
	c := cron.New(
		cron.WithLogger(cron.DefaultLogger),
	)

	cronLoopSecond := env.GetString("CRON_EVER_SECOND", "5")
	_, err = c.AddFunc(fmt.Sprintf("@every %ss", cronLoopSecond), func() {
		printLN := fmt.Sprintf("===== Cron Started on: %s ====", time.Now().Add(time.Hour*7))
		fmt.Println(printLN)

		transactions, err := uc.db.GetTransactionByCode(ctx)
		if err != nil {
			fmt.Println(err)
		}
		if len(transactions) > 0 {
			go func() {
				for _, flightTransaction := range transactions {
					for _, flightSegment := range flightTransaction.FlightTransaction {
						_, err = uc.flightProviders.FlightCancelBooking(ctx, models.RequestAfterBook{
							PnrId:     flightSegment.PnrCode,
							BookingId: flightSegment.BookingCode,
							UserId:    "",
						})
						if err != nil {
							fmt.Printf("error: %s \n", err.Error())
						}
					}
				}
			}()

			err = uc.db.UpdateTransactionToExpired(ctx)
			if err != nil {
				logger.LogError(err)
			}
		}
	})
	if err != nil {
		logger.LogError(err)
	}

	c.Start()

	return nil
}
