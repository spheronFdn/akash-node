package inflation

import (
	types "github.com/spheronFdn/akash-api-fork/go/node/inflation/v1beta3"

	"github.com/spheronFdn/akash-node/x/inflation/keeper"
)

const (
	// StoreKey represents storekey of inflation module
	StoreKey = types.StoreKey
	// ModuleName represents current module name
	ModuleName = types.ModuleName
)

type (
	// Keeper defines keeper of inflation module
	Keeper = keeper.IKeeper
)

var (
	// NewKeeper creates new keeper instance of inflation module
	NewKeeper = keeper.NewKeeper
)
