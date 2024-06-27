// Package v0_26_0
// nolint revive
package v0_26_0

import (
	utypes "github.com/spheronFdn/akash-node/upgrades/types"
)

func init() {
	utypes.RegisterUpgrade(UpgradeName, initUpgrade)
}
