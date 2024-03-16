package db

import (
	"github.com/carpetsage/EggContractor/api"
	"github.com/carpetsage/EggContractor/coop"
)

// Contract is set to nil if no matching contract is found or if there is a
// database error.
func WrapCoopStatusWithDB(c *api.CoopStatus) (*coop.CoopStatus, error) {
	contract, err := GetContract(c.ContractId, 0, 0)
	return &coop.CoopStatus{
		CoopStatus: c,
		Contract:   contract,
	}, err
}
