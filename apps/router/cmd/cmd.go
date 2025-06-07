package cmd

import "gitlab.com/tixia-backend/go-boilerplate/apps/service"

type cmd struct {
	transaction service.ITransaction
}

func NewCmd(
	transaction service.ITransaction,
) *cmd {
	return &cmd{
		transaction: transaction,
	}
}
