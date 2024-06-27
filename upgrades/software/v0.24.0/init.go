// Package v0_24_0
// nolint revive
package v0_24_0

import (
	dv1beta3 "github.com/spheronFdn/akash-api-fork/go/node/deployment/v1beta3"
	mv1beta3 "github.com/spheronFdn/akash-api-fork/go/node/market/v1beta3"

	utypes "github.com/spheronFdn/akash-node/upgrades/types"
)

func init() {
	utypes.RegisterUpgrade(upgradeName, initUpgrade)

	utypes.RegisterMigration(dv1beta3.ModuleName, 2, newDeploymentMigration)
	utypes.RegisterMigration(mv1beta3.ModuleName, 2, newMarketMigration)
}
