package market

import (
	types "github.com/spheronFdn/akash-api-fork/go/node/market/v1beta4"

	"github.com/spheronFdn/akash-node/x/market/keeper"
)

const (
	// StoreKey represents storekey of market module
	StoreKey = types.StoreKey
	// ModuleName represents current module name
	ModuleName = types.ModuleName
)

type (
	// Keeper defines keeper of market module
	Keeper = keeper.Keeper
)

var (
	// NewKeeper creates new keeper instance of market module
	NewKeeper = keeper.NewKeeper
)
