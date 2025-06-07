package cmd

import (
	"context"

	"github.com/vizucode/gokit/logger"
)

func (h *cmd) DoExpired() (err error) {
	ctx := context.Background()

	err = h.transaction.DoExpired(ctx)
	if err != nil {
		logger.Log.Debug(ctx)
	}

	return nil
}
