package flightproviders

import (
	"context"
	"net/http"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gitlab.com/tixia-backend/go-boilerplate/apps/models"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/httpstd"
	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/rpcstd"
	pb "gitlab.com/tixia-backend/go-boilerplate/proto"
)

func (repo *flightProvider) FlightCancelBooking(ctx context.Context, req models.RequestAfterBook) (resp models.RequestAfterBook, err error) {

	var (
		payloadFlightCancel = &pb.RequestAfterBook{
			PnrId:     req.PnrId,
			UserId:    req.UserId,
			BookingId: req.BookingId,
		}
	)

	resultCancelBook, err := repo.client.FlightCancelBooking(ctx, payloadFlightCancel)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, httpstd.InternalServerError)
	}

	resp = models.RequestAfterBook{
		PnrId:  resultCancelBook.GetPnrId(),
		UserId: resultCancelBook.GetUserId(),
	}

	return resp, nil
}
