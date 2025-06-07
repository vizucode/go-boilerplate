package repositories

import (
	"context"

	"gitlab.com/tixia-backend/go-boilerplate/apps/models"
)

type IFlightProviders interface {
	FlightCancelBooking(ctx context.Context, req models.RequestAfterBook) (resp models.RequestAfterBook, err error)
}

type IDatabase interface {
	UpdateTransactionToExpired(ctx context.Context) (err error)
	GetTransactionByCode(ctx context.Context) (resp []models.TbTransaction, err error)
}
