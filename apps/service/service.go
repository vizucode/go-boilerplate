package service

import (
	"context"
)

type ITransaction interface {
	DoExpired(ctx context.Context) (err error)
}
