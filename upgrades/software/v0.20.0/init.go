// Package v0_20_0
// nolint revive
package v0_20_0

import (
	utypes "github.com/spheronFdn/akash-node/upgrades/types"
)

func init() {
	utypes.RegisterUpgrade(upgradeName, initUpgrade)
}
