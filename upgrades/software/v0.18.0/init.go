// Package v0_18_0
// nolint revive
package v0_18_0

import (
	utypes "github.com/spheronFdn/akash-node/upgrades/types"
)

func init() {
	utypes.RegisterUpgrade(upgradeName, initUpgrade)
}
