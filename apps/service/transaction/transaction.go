package transaction

import "gitlab.com/tixia-backend/go-boilerplate/apps/repositories"

type transaction struct {
	db              repositories.IDatabase
	flightProviders repositories.IFlightProviders
}

func NewTransaction(
	db repositories.IDatabase,
	flightProviders repositories.IFlightProviders,
) *transaction {
	return &transaction{
		db:              db,
		flightProviders: flightProviders,
	}
}
