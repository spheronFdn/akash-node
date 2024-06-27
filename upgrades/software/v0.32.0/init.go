// Package v0_32_0
// nolint revive
package v0_32_0

import (
	mv1beta4 "github.com/spheronFdn/akash-api-fork/go/node/market/v1beta4"

	utypes "github.com/spheronFdn/akash-node/upgrades/types"
)

func init() {
	utypes.RegisterUpgrade(UpgradeName, initUpgrade)
	utypes.RegisterMigration(mv1beta4.ModuleName, 4, newMarketMigration)
}
