// Package v0_30_0
// nolint revive
package v0_30_0

import (
	utypes "github.com/spheronFdn/akash-node/upgrades/types"
)

func init() {
	utypes.RegisterUpgrade(UpgradeName, initUpgrade)
}
